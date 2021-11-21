package main

import "github.com/joho/godotenv"

func main() {
	loadEnv()

	InitDatabase()

	api := Api{}
	api.Initialize()

	api.Run()
}

func loadEnv() {
	godotenv.Load()
}
