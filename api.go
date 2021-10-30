package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
}

func (api *Api) Initialize() {
	api.Router = mux.NewRouter()

	// tom: this line is added after initializeRoutes is created later on
	api.createRoutes()
}

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

func (api *Api) respondWithError(w http.ResponseWriter, code int, message string) {
	api.respondWithJSON(w, code, map[string]string{"error": message})
}

func (api *Api) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Handler functions
func (api *Api) search(w http.ResponseWriter, r *http.Request) {
	api.respondWithJSON(w, 200, map[string]bool{"found": true})
}

func (api *Api) createRoutes() {
	api.Router.HandleFunc("/search", api.search).Methods("POST")
}
