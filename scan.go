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

	numberOfPackages := len(packageList)

	// Keep track of the packages loaded to later make the summary
	ch := make(chan Package, numberOfPackages)

	// We use a counter here to keep track of the current index.
	// It is then used to assign a loaded package to the packages slice from above
	for name, version := range packageList {
		wg.Add(1)

		// load every package, put it into the package list and send the response via websocket
		go func(name string, version string, packageManager PackageManger, conn *websocket.Conn, wg *sync.WaitGroup) {
			defer wg.Done()

			pkg, err := loadPackage(packageManager, name)

			var response PackageScanMessage

			if err != nil {
				response = constructScanPackageResponse(pkg, name, version, false)
			} else {
				response = constructScanPackageResponse(pkg, name, version, true)

				ch <- pkg
			}

			sendScanPackageResponse(response, conn)
		}(name, version, packageManager, conn, &wg)
	}

	wg.Wait()

	close(ch)

	var packages []Package

	for pkg := range ch {
		packages = append(packages, pkg)
	}

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

// constructSummary uses the conditions database to generate a SummaryScanMessage for the given packages.
func constructSummary(packages []Package) SummaryScanMessage {
	var licenseIds []string

	for _, pkg := range packages {
		licenseIds = append(licenseIds, pkg.License.License)
	}

	details, err := DB.GetLicenseConditionsByIds(licenseIds)

	if err != nil {
		log.Fatal("Failed to query data for summary.")
	}

	conditions := ConditionsToArray(details)

	// Now the actual "business logic": For every row that has no "false" in it, add the condition/permission/limitation to the list.
	// We do this in a manual way to make it easy to add more complex logic in the future.

	var summary SummaryConditions

	// Permissions
	if !boolArrayContains(conditions["CommercialUse"], false) {
		summary.Permissions = append(summary.Permissions, "Commercial use")
	}
	if !boolArrayContains(conditions["Distribution"], false) {
		summary.Permissions = append(summary.Conditions, "Distribution")
	}
	if !boolArrayContains(conditions["Modification"], false) {
		summary.Permissions = append(summary.Permissions, "Modification")
	}
	if !boolArrayContains(conditions["PatentUse"], false) {
		summary.Permissions = append(summary.Conditions, "Patent use")
	}
	if !boolArrayContains(conditions["PrivateUse"], false) {
		summary.Permissions = append(summary.Permissions, "Private use")
	}

	// Conditions
	if boolArrayContains(conditions["DiscloseSource"], true) {
		summary.Conditions = append(summary.Conditions, "Disclose source")
	}
	if boolArrayContains(conditions["LicenseAndCopyrightNotice"], true) {
		summary.Conditions = append(summary.Conditions, "License and copyright notice")
	}
	if boolArrayContains(conditions["LicenseAndCopyrightNoBinaries"], true) {
		summary.Conditions = append(summary.Conditions, "License and copyright notice for binaries")
	}
	if boolArrayContains(conditions["NetworkUseIsDistribution"], true) {
		summary.Conditions = append(summary.Conditions, "Network use is distribution")
	}
	if boolArrayContains(conditions["SameLicense"], true) {
		summary.Conditions = append(summary.Conditions, "Same license")
	}
	if boolArrayContains(conditions["StateChanges"], true) {
		summary.Conditions = append(summary.Conditions, "State changes")
	}

	// Limitations
	if boolArrayContains(conditions["Liability"], true) {
		summary.Conditions = append(summary.Conditions, "Limited Liability")
	}
	if boolArrayContains(conditions["TrademarkUse"], true) {
		summary.Conditions = append(summary.Conditions, "No trademark use")
	}
	if boolArrayContains(conditions["Warranty"], true) {
		summary.Conditions = append(summary.Conditions, "No warranty")
	}

	response := SummaryScanMessage{
		Type: SummaryMessage,
		Data: Summary{
			Conditions: summary,
		},
	}

	return response
}

func boolArrayContains(s []bool, e bool) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
