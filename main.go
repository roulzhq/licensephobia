package main

import "github.com/joho/godotenv"

type App struct {
	api Api
	db  Database
}

func main() {
	loadEnv()

	db := Database{}
	api := Api{}

	app := App{api, db}

	app.InitDb()
	app.InitApi()
	app.RunApi()
}

func loadEnv() {
	godotenv.Load()
}
