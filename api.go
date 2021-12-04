package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/roulzhq/licensephobia/database"
)

type Api struct {
	Router   *mux.Router
	Upgrader websocket.Upgrader
}

type Ping struct {
	Ping string `json:"ping"`
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

type SearchPreviewRequest struct {
	PackageManager PackageManger `json:"packageManager"`
	Name           string        `json:"name"`
}

type SearchRequest struct {
	PackageManager PackageManger `json:"packageManager"`
	Name           string        `json:"name"`
}

// Initialize creates the API router and route endpoints
func (api *Api) Init() {
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
func (api *Api) Run(port int) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"https://api.licensephobia.com",
			"http://api.licensephobia.com",
			"https://dev.api.licensephobia.com",
			"http://dev.api.licensephobia.com",
		},
		AllowCredentials: true,
	})

	handler := c.Handler(api.Router)
	server := &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	log.Println("Running server at ", server.Addr)

	log.Fatal(server.ListenAndServe())
}

func (api *Api) createRoutes() {
	api.Router.Path("/scan").HandlerFunc(api.scan)
	api.Router.Path("/searchPreview").HandlerFunc(api.searchPreview)
	api.Router.Path("/search").Queries("packageManager", "{packageManager}", "name", "{name}").HandlerFunc(api.search)

	api.Router.Path("/licenses").HandlerFunc(api.getLicenses)
	api.Router.Path("/licenses/conditions").Queries("id", "{id}").HandlerFunc(api.getLicenseConditions)
	api.Router.Path("/licenses/conditions").HandlerFunc(api.getLicenseConditions)
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

func (api *Api) getLicenses(w http.ResponseWriter, r *http.Request) {
	licenses, err := DB.GetLicenses()

	if err != nil {
		log.Println(err.Error())
		api.respondWithError(w, 500, "Unable to load licenses from database")
		return
	}

	api.respondWithJSON(w, 200, licenses)
}

func (api *Api) getLicenseConditions(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	ids := params["id"]

	var licenses []database.LicenseConditions
	var err error

	if len(ids) > 0 {
		licenses, err = DB.GetLicenseConditionsByIds(ids)
	} else {
		licenses, err = DB.GetLicenseConditions()
	}

	if err != nil {
		log.Println(err.Error())
		api.respondWithError(w, 500, "Unable to load licenses from database")
		return
	}

	api.respondWithJSON(w, 200, licenses)
}

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

func (api *Api) searchPreview(w http.ResponseWriter, r *http.Request) {
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

		var searchRequest SearchPreviewRequest
		json.Unmarshal(message, &searchRequest)

		HandleSearchPreviewRequest(searchRequest, conn)
	}
}

func (api *Api) search(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	packageManager := params.Get("packageManager")
	name := params.Get("name")

	request := SearchRequest{PackageManger(packageManager), name}

	response, err := HandleSearchRequest(request)

	if err != nil {
		api.respondWithError(w, 404, "Could not find the package you where looking for.")
	} else {
		api.respondWithJSON(w, 200, response)
	}
}
