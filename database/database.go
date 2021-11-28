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

func (db *Database) GetLicenses() ([]map[string]interface{}, error) {
	res, err := db.Client.From("licenses").Select("*", "", false).Execute()

	if err != nil {
		return nil, err
	}

	var responseMap []map[string]interface{}

	err = json.Unmarshal(res, &responseMap)

	if err != nil {
		log.Fatal(err)
	}

	return responseMap, nil
}

func (db *Database) GetLicenseById(id string) (map[string]interface{}, error) {
	res, err := db.Client.From("licenses").Select("*", "", false).Eq("licenseId", id).Execute()

	if err != nil {
		return nil, err
	}

	var responseMap []map[string]interface{}

	_ = json.Unmarshal(res, &responseMap)

	if len(responseMap) == 0 {
		return nil, errors.New("License not found")
	}

	return responseMap[0], nil
}

func (db *Database) GetLicenseNameById(id string) (map[string]interface{}, error) {
	res, err := db.Client.From("licenses").Select("name", "", false).Eq("licenseId", id).Execute()

	if err != nil {
		return nil, err
	}

	var responseMap []map[string]interface{}

	_ = json.Unmarshal(res, &responseMap)

	if len(responseMap) == 0 {
		return nil, errors.New("License not found")
	}

	return responseMap[0], nil
}

func (db *Database) GetLicenseConditions() ([]map[string]interface{}, error) {
	res, err := db.Client.From("conditions").Select("*", "", false).Execute()

	if err != nil {
		return nil, err
	}

	var responseMap []map[string]interface{}

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
