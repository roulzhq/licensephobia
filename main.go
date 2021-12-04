package main

import (
	"flag"

	"github.com/joho/godotenv"
	db "github.com/roulzhq/licensephobia/database"
)

var DB db.Database = db.Database{}

func main() {
	godotenv.Load()

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
	api := Api{}

	DB.Init()
	api.Init()
	api.Run(port)
}

func runDataImporter() {
	db.ImportLicenseData()
}
