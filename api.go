package main

import (
	"encoding/json"
	"log"
	"net/http"
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

// Initialize creates the API router and route endpoints
func (api *Api) Initialize() {
	api.Router = mux.NewRouter()
	api.Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:3000"
		},
	}

	// tom: this line is added after initializeRoutes is created later on
	api.createRoutes()
}

// Run serves the API via a webserver
func (api *Api) Run() {
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Router,
	}

	log.Println("Running server at ", server.Addr)

	log.Fatal(server.ListenAndServe())
}

func (api *Api) createRoutes() {
	api.Router.HandleFunc("/scan", api.scan)
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

func (api *Api) scan(w http.ResponseWriter, r *http.Request) {
	conn, err := api.Upgrader.Upgrade(w, r, nil)

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
