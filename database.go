package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/supabase/postgrest-go"
)

type Database struct {
	Client *postgrest.Client
}

func (app *App) InitDb() {
	host := os.Getenv("SUPABASE_HOST")
	key := os.Getenv("SUPABASE_KEY")

	url := "https://" + host

	app.db.Client = postgrest.NewClient(url, "public", map[string]string{})

	app.db.Client.TokenAuth(key)

	if app.db.Client.ClientError != nil {
		panic(app.db.Client.ClientError)
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
