package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/supabase/postgrest-go"
)

type Database struct {
	Client *postgrest.Client
}

func InitDatabase() {
	host := os.Getenv("SUPABASE_HOST")
	key := os.Getenv("SUPABASE_KEY")

	url := "https://" + host

	c := postgrest.NewClient(url, "public", map[string]string{})

	c = c.TokenAuth(key)

	res, err := c.From("licenses").Select("license", "", false).Execute()

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func (db *Database) Connect() {
	password := os.Getenv("SUPABASE_PASSWORD")
	user := os.Getenv("SUPABASE_USER")
	host := os.Getenv("SUPABASE_HOST")
	port := os.Getenv("SUPABASE_PORT")
	dbName := os.Getenv("SUPABASE_DB")

	connUrl := "postgresql://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName
	connUrl = url.QueryEscape(connUrl)

	db.Client = postgrest.NewClient(connUrl, "public", map[string]string{})

	if db.Client.ClientError != nil {
		panic(db.Client.ClientError)
	}

	fmt.Println(db.Client)
}
