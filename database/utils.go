package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type LicenseDataJson struct {
	LicenseListVersion string            `json:"licenseListVersion"`
	Licenses           []LicenseDataItem `json:"licenses"`
}

type LicenseDataItem struct {
	Reference             string   `json:"reference"`
	IsDeprecatedLicenseId bool     `json:"isDeprecatedLicenseId"`
	DetailsUrl            string   `json:"detailsUrl"`
	ReferenceNumber       int      `json:"referenceNumber"`
	Name                  string   `json:"name"`
	LicenseId             string   `json:"licenseId"`
	SeeAlso               []string `json:"seeAlso"`
	IsOsiApproved         bool     `json:"isOsiApproved"`
	IsFsfLibre            bool     `json:"isFsfLibre"`
}

func ImportLicenseData() {
	db := Database{}
	db.Init()

	licenseDataUrl := "https://raw.githubusercontent.com/spdx/license-list-data/master/json/licenses.json"

	response, err := http.Get(licenseDataUrl)

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var licenseData LicenseDataJson
	err = json.Unmarshal(responseData, &licenseData)

	if err != nil {
		log.Fatal(err)
	}

	db.UpsertLicenses(licenseData.Licenses)
}
