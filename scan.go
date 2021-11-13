package main

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type ScanResponse struct {
	Id          string              `json:"id"`
	Found       bool                `json:"found"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Version     ScanResponseVersion `json:"version"`
	License     ScanResponseLicense `json:"license"`
	Url         string              `json:"url"`
}

type ScanResponseVersion struct {
	Used   string `json:"used"`
	Latest string `json:"latest"`
}

type ScanResponseLicense struct {
	Found       bool   `json:"found"`
	LicenseType string `json:"type"`
}

func HandleScanRequest(scanRequest ScanRequest, conn *websocket.Conn) error {
	packageManager := scanRequest.PackageManager
	mimeType, file, err := decodeFileString(scanRequest.Data)

	if err != nil {
		log.Println(err)
		return err
	}

	switch packageManager {
	case "npm":
		if mimeType == "data:application/json;" {
			ScanPackageJson(file, conn)
		} else {
			return errors.New("the file you uploaded is not a valid package.json file")
		}
	case "pip":
		break
	case "cargo":
		break
	}

	return nil
}

func HandleSearchRequest(searchRequest SearchRequest, conn *websocket.Conn) error {
	packageManager := searchRequest.PackageManager

	switch packageManager {
	case "npm":
		SearchPackage(searchRequest.Data, conn)
	case "pip":
		break
	case "cargo":
		break
	}

	return nil
}

func decodeFileString(file string) (mimeType string, data []byte, err error) {
	splitString := strings.Split(string(file), "base64,")
	mimeType = splitString[0]
	b64data := splitString[1]

	dec, err := base64.StdEncoding.DecodeString(b64data)

	data = dec

	return mimeType, data, err
}
