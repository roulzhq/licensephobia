package database

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/supabase/postgrest-go"
)

type Database struct {
	Client *postgrest.Client
}

func (db *Database) Init() {
	host := os.Getenv("SUPABASE_HOST")
	key := os.Getenv("SUPABASE_KEY")

	url := "https://" + host

	db.Client = postgrest.NewClient(url, "public", map[string]string{})

	db.Client.TokenAuth(key)

	if db.Client.ClientError != nil {
		panic(db.Client.ClientError)
	}
}

func (db *Database) GetLicenses() ([]License, error) {
	res, err := db.Client.From("licenses").Select("*", "", false).Execute()

	if err != nil {
		return nil, err
	}

	var response []License

	err = json.Unmarshal(res, &response)

	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}

func (db *Database) GetLicenseById(id string) (License, error) {
	res, err := db.Client.From("licenses").Select("*", "", false).Eq("licenseId", id).Execute()

	if err != nil {
		return License{}, err
	}

	var response []License

	_ = json.Unmarshal(res, &response)

	if len(response) == 0 {
		return License{}, errors.New("License not found")
	}

	return response[0], nil
}

func (db *Database) GetLicenseNameById(id string) (string, error) {
	res, err := db.Client.From("licenses").Select("name", "", false).Eq("licenseId", id).Execute()

	if err != nil {
		return "", err
	}

	var responseMap []string

	_ = json.Unmarshal(res, &responseMap)

	if len(responseMap) == 0 {
		return "", errors.New("License not found")
	}

	return responseMap[0], nil
}

func (db *Database) GetLicenseConditions() ([]LicenseConditions, error) {
	res, err := db.Client.From("conditions").Select("*", "", false).Execute()

	if err != nil {
		return nil, err
	}

	var responseMap []LicenseConditions

	err = json.Unmarshal(res, &responseMap)

	if err != nil {
		log.Fatal(err)
	}

	return responseMap, nil
}

func (db *Database) GetLicenseConditionsByIds(licenseIds []string) ([]LicenseConditions, error) {
	res, err := db.Client.From("conditions").Select("*", "", false).In("licenseId", licenseIds).Execute()

	if err != nil {
		return nil, err
	}

	var responseMap []LicenseConditions

	err = json.Unmarshal(res, &responseMap)

	if err != nil {
		log.Fatal(err)
	}

	return responseMap, nil
}

func (db *Database) UpsertLicenses(licenses []LicenseDataItem) ([]byte, error) {
	res, err := db.Client.From("licenses").Insert(licenses, true, "", "", "").Execute()

	if err != nil {
		return nil, err
	}

	return res, nil
}
