package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "github.com/roulzhq/licensephobia/database"
)

type App struct {
	api Api
	db  db.Database
}

func main() {
	godotenv.Load()

	argsWithProg := os.Args

	log.Print(argsWithProg)

	port := flag.Int("p", 8080, "port to use for the api")
	importLicenseData := flag.Bool("importLicenseData", false, "Import license data from the SPDX github repo, this will NOT run the api")

	flag.Parse()

	if *importLicenseData {
		runDataImporter()
	} else {
		runApp(*port)
	}
}

func runApp(port int) {
	db := db.Database{}
	api := Api{}

	app := App{api, db}

	app.db.Init()
	app.InitApi()
	app.RunApi(port)
}

func runDataImporter() {
	db.ImportLicenseData()
}
