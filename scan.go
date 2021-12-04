package main

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"
	"sync"

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
	var wg sync.WaitGroup
	var packageList map[string]string

	defer wg.Done()
	defer conn.Close()

	packageManager := scanRequest.PackageManager
	mimeType, file, err := decodeFileString(scanRequest.Data)

	if err != nil {
		log.Println(err)
		return err
	}

	switch packageManager {
	case "npm":
		if mimeType == "data:application/json;" {
			packageList = ScanPackageJson(file)
		} else {
			return errors.New("the file you uploaded is not a valid package.json file")
		}
	case "pip":
		break
	case "cargo":
		break
	}

	// Keep track of the packages loaded to later make the summary
	packages := make(map[string]*Package, len(packageList))
	var packagesMutex = sync.RWMutex{}

	// We use a counter here to keep track of the current index.
	// It is then used to assign a loaded package to the packages slice from above
	for name, version := range packageList {
		wg.Add(1)

		// load every package, put it into the package list and send the response via websocket
		go func(name string, version string, packageManager PackageManger, conn *websocket.Conn, wg *sync.WaitGroup, packages map[string]*Package) {
			defer wg.Done()

			pkg, err := loadPackage(packageManager, name)

			var response PackageScanMessage

			if err != nil {
				response = constructScanPackageResponse(pkg, name, version, false)
			} else {
				response = constructScanPackageResponse(pkg, name, version, true)

				packagesMutex.RLock()
				packages[name] = &pkg
				packagesMutex.RUnlock()
			}

			sendScanPackageResponse(response, conn)
		}(name, version, packageManager, conn, &wg, packages)
	}

	wg.Wait()

	// When all packages are loaded, construct the summary
	summary := constructSummary(packages)
	sendScanSummaryResponse(summary, conn)

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

func constructScanPackageResponse(pkg Package, usedName string, usedVersion string, found bool) PackageScanMessage {
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

func sendScanPackageResponse(response PackageScanMessage, conn *websocket.Conn) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	conn.WriteJSON(response)
}

func sendScanSummaryResponse(response SummaryScanMessage, conn *websocket.Conn) {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()
	conn.WriteJSON(response)
}

func constructSummary(packages map[string]*Package) SummaryScanMessage {
	response := SummaryScanMessage{
		Type: SummaryMessage,
		Data: Summary{
			Conditions: SummaryConditions{
				Permissions: []string{"Commercial Use", "Distribution", "Modification", "Patent use", "Private use"},
				Conditions:  []string{"Disclose source", "License notice", "Copyright notice", "Network use is distribution", "Same license", "State changes"},
				Limitations: []string{"Liability", "Warranty"},
			},
		},
	}

	return response
}

func loadPackage(packageManager PackageManger, name string) (Package, error) {
	switch packageManager {
	case "npm":
		return GetNpmPackage(name)
	case "pip":
		break
	case "cargo":
		break
	default:
		return Package{}, errors.New("the given packageManager is not supported")
	}

	return Package{}, errors.New("unknown error while getting the package")
}
