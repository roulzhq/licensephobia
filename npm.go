package main

import (
	"encoding/json"
	"errors"
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

func GetNpmPackage(name string) (Package, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "registry.npmjs.org",
		Path:   name,
	}

	response, err := http.Get(url.String())

	if response.StatusCode != http.StatusOK || err != nil {
		return Package{}, errors.New("Could not find NPM package")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err.Error())
		return Package{}, err
	}

	var responseJson ApiPackageResponse
	json.Unmarshal(responseData, &responseJson)

	pkg := NpmPackageToGeneric(responseJson)

	return pkg, nil
}

func NpmPackageToGeneric(npmPackage ApiPackageResponse) Package {
	known := LicenseExists(npmPackage.License)

	return Package{
		Id:            npmPackage.Id,
		Name:          npmPackage.Name,
		Description:   npmPackage.Description,
		Homepage:      npmPackage.Homepage,
		LatestVersion: npmPackage.DistTags["latest"],
		License: LicenseInfo{
			Found:   npmPackage.License != "",
			Known:   known,
			License: npmPackage.License,
		},
	}
}

func sendPackageResponse(name string, version string, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	responseJson, err := GetNpmPackage(name)

	var socketData PackageScanMessage

	if err != nil {
		socketData = ConstructScanPackageResponse(responseJson, name, version, false)
	} else {
		socketData = ConstructScanPackageResponse(responseJson, name, version, true)
	}

	SendScanPackageResponse(socketData, conn)
}
