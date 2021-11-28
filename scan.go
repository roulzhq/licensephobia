package main

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

type ScanMessageType string

const (
	PackageMessage ScanMessageType = "package"
	SummaryMessage ScanMessageType = "summary"
)

type PackageScanMessage struct {
	Type ScanMessageType `json:"type"`
	Data ScanResponse    `json:"data"`
}

type SummaryScanMessage struct {
	Type ScanMessageType `json:"type"`
	Data Summary         `json:"data"`
}

type ScanResponse struct {
	Found   bool    `json:"found"`
	Name    string  `json:"name"`
	Version string  `json:"version"`
	Package Package `json:"package"`
}

type SummaryResponse = Summary

type ScanResponseVersion struct {
	Used   string `json:"used"`
	Latest string `json:"latest"`
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

func decodeFileString(file string) (mimeType string, data []byte, err error) {
	splitString := strings.Split(string(file), "base64,")
	mimeType = splitString[0]
	b64data := splitString[1]

	dec, err := base64.StdEncoding.DecodeString(b64data)

	data = dec

	return mimeType, data, err
}

// Logic to send package data via websockets

func ConstructScanPackageResponse(pkg Package, usedName string, usedVersion string, found bool) PackageScanMessage {
	scanResponse := ScanResponse{
		Name:    usedName,
		Version: usedVersion,
		Found:   found,
		Package: pkg,
	}

	return PackageScanMessage{
		Type: PackageMessage,
		Data: scanResponse,
	}
}

func SendScanPackageResponse(response PackageScanMessage, conn *websocket.Conn) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	conn.WriteJSON(response)
}

func SendScanSummaryResponse(response SummaryScanMessage, conn *websocket.Conn) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	conn.WriteJSON(response)
}
