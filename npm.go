package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

type Author struct {
	Name string `json:"name"`
}

type ApiPackageResponse struct {
	Id          string            `json:"_id"`
	Name        string            `json:"name"`
	DistTags    map[string]string `json:"dist-tags"`
	License     string            `json:"license"`
	Author      Author            `json:"author"`
	Homepage    string            `json:"homepage"`
	Description string            `json:"description"`
}

type PackageJson struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

var connectionMutex sync.Mutex

func SearchNpmPackage(name string) []SearchPreviewResponse {
	v := url.Values{}

	v.Set("q", name)

	url := url.URL{
		Scheme:   "https",
		Host:     "www.npmjs.com",
		Path:     "search/suggestions",
		RawQuery: v.Encode(),
	}

	response, err := http.Get(url.String())

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseJson []SearchPreviewResponse
	json.Unmarshal(responseData, &responseJson)

	if err != nil {
		log.Fatal(err)
	}

	if len(responseJson) > 5 {
		responseJson = responseJson[:5]
	}

	return responseJson
}

func ScanPackageJson(file []byte, conn *websocket.Conn) {
	var packageJson PackageJson

	json.Unmarshal(file, &packageJson)

	var wg sync.WaitGroup

	for name, version := range packageJson.Dependencies {
		wg.Add(1)
		go sendPackageResponse(name, version, conn, &wg)
	}

	for name, version := range packageJson.DevDependencies {
		wg.Add(1)
		go sendPackageResponse(name, version, conn, &wg)
	}

	wg.Wait()
	conn.Close()
}

func constructPackageResponse(npmPackage ApiPackageResponse, usedVersion string) ScanResponse {
	return ScanResponse{
		Id:          npmPackage.Id,
		Name:        npmPackage.Name,
		Found:       true,
		Description: npmPackage.Description,
		Url:         npmPackage.Homepage,
		Version: ScanResponseVersion{
			Used:   usedVersion,
			Latest: npmPackage.DistTags["latest"],
		},
		License: ScanResponseLicense{
			Found:       npmPackage.License != "",
			LicenseType: npmPackage.License,
		},
	}
}

func sendResponseToSocket(data ScanResponse, conn *websocket.Conn) {
	// TODO: Check if using a mutex like this actually protects the socket
	// https://github.com/gorilla/websocket/issues/119
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	conn.WriteJSON(data)
}

func sendPackageResponse(name string, version string, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	responseJson := GetNpmPackage(name)

	socketData := constructPackageResponse(responseJson, version)

	sendResponseToSocket(socketData, conn)
}

func GetNpmPackage(name string) ApiPackageResponse {
	url := url.URL{
		Scheme: "https",
		Host:   "registry.npmjs.org",
		Path:   name,
	}

	response, err := http.Get(url.String())

	if err != nil {
		log.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseJson ApiPackageResponse
	json.Unmarshal(responseData, &responseJson)

	return responseJson
}
