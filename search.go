package main

import (
	"errors"

	"github.com/gorilla/websocket"
)

type SearchResponse struct {
	Id          string              `json:"id"`
	Found       bool                `json:"found"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Version     string              `json:"version"`
	License     ScanResponseLicense `json:"license"`
	Url         string              `json:"url"`
}

type SearchPreviewResponse struct {
	Name string `json:"name"`
}

func HandleSearchPreviewRequest(search SearchPreviewRequest, conn *websocket.Conn) error {
	packageManager := search.PackageManager

	var response []SearchPreviewResponse

	switch packageManager {
	case "npm":
		response = SearchNpmPackage(search.Name)
	case "pip":
		break
	case "cargo":
		break
	}

	conn.WriteJSON(response)

	return nil
}

func HandleSearchRequest(search SearchRequest) (SearchResponse, error) {
	packageManager := search.PackageManager

	var response SearchResponse
	var error error = nil

	switch packageManager {
	case "npm":
		npmPackage := GetNpmPackage(search.Name)
		response = apiToSearchResponse(npmPackage)
	case "pip":
		break
	case "cargo":
		break
	}

	if error != nil {
		return response, errors.New("could not find package")
	} else {
		return response, nil
	}
}

func apiToSearchResponse(npmPackage ApiPackageResponse) SearchResponse {
	return SearchResponse{
		Id:          npmPackage.Id,
		Name:        npmPackage.Name,
		Found:       true,
		Description: npmPackage.Description,
		Url:         npmPackage.Homepage,
		Version:     npmPackage.DistTags["latest"],
		License: ScanResponseLicense{
			Found:       npmPackage.License != "",
			LicenseType: npmPackage.License,
		},
	}
}
