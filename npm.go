package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

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

func ScanPackageJson(file []byte, conn *websocket.Conn) {
	log.Print(file)

	var packageJson PackageJson

	json.Unmarshal(file, &packageJson)

	log.Println(packageJson)

	for name, version := range packageJson.Dependencies {
		go sendPackageResponse(name, version, conn)
	}

	for name, version := range packageJson.DevDependencies {
		go sendPackageResponse(name, version, conn)
	}
}

func constructScanPackageResponse(npmPackage ApiPackageResponse, usedVersion string) ScanResponse {
	return ScanResponse{
		Id:          npmPackage.Id,
		Name:        npmPackage.Name,
		Found:       true,
		Description: npmPackage.Description,
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

func sendPackageResponse(name string, version string, conn *websocket.Conn) {
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

	socketData := constructScanPackageResponse(responseJson, version)

	conn.WriteJSON(socketData)
}
