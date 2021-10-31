package main

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"

	"github.com/gorilla/websocket"

	"github.com/roulzhq/licensephobia/internal/npm"
)

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
			npm.ScanPackageJson(file, conn)
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
