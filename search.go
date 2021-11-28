package main

import (
	"errors"

	"github.com/gorilla/websocket"
)

type SearchResponse struct {
	Found   bool    `json:"found"`
	Package Package `json:"package"`
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
		pkg, err := GetNpmPackage(search.Name)

		if err != nil {
			response = apiToSearchResponse(pkg, false)
		} else {
			response = apiToSearchResponse(pkg, true)
		}
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

func apiToSearchResponse(pkg Package, found bool) SearchResponse {
	return SearchResponse{
		Found:   found,
		Package: pkg,
	}
}
