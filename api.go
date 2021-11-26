package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Api struct {
	Router   *mux.Router
	Upgrader websocket.Upgrader
}

type PackageManger string

const (
	Npm   PackageManger = "npm"
	Pip   PackageManger = "pip"
	Cargo PackageManger = "cargo"
)

type ScanRequest struct {
	PackageManager PackageManger `json:"packageManager"`
	Data           string        `json:"data"`
}

type SearchRequest struct {
	PackageManager PackageManger `json:"packageManager"`
	Data           string        `json:"data"`
}

// Initialize creates the API router and route endpoints
func (app *App) InitApi() {
	app.api.Router = mux.NewRouter()
	app.api.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:3000"
		},
	}

	// tom: this line is added after initializeRoutes is created later on
	app.createRoutes()
}

// Run serves the API via a webserver
func (app *App) RunApi(port int) {
	if port == 0 {
		port = 8080
	}

	server := &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      app.api.Router,
	}

	log.Println("Running server at ", server.Addr)

	log.Fatal(server.ListenAndServe())
}

func (app *App) createRoutes() {
	app.api.Router.HandleFunc("/scan", app.scan)
	app.api.Router.HandleFunc("/search", app.search)

	app.api.Router.HandleFunc("/licenses", app.getLicenses)
}

func (api *Api) respondWithError(w http.ResponseWriter, code int, message string) {
	api.respondWithJSON(w, code, map[string]string{"error": message})
}

func (api *Api) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// ------------------------------------------
// Handler functions
// ------------------------------------------

func (app *App) getLicenses(w http.ResponseWriter, r *http.Request) {
	licenses, err := app.db.GetLicenses()

	if err != nil {
		log.Println(err.Error())
		app.api.respondWithError(w, 500, "Unable to load licenses from database")
		return
	}

	app.api.respondWithJSON(w, 200, licenses)
}

func (app *App) scan(w http.ResponseWriter, r *http.Request) {
	conn, err := app.api.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var scanRequest ScanRequest
		json.Unmarshal(message, &scanRequest)

		HandleScanRequest(scanRequest, conn)
	}
}

func (app *App) search(w http.ResponseWriter, r *http.Request) {
	conn, err := app.api.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var searchRequest SearchRequest
		json.Unmarshal(message, &searchRequest)

		HandleSearchRequest(searchRequest, conn)
	}
}
