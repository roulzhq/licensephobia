package npm

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Author struct {
	Name string `json:"name"`
}

type ApiPackageResponse struct {
	Id          string            `json:"_id"`
	Name        string            `json:"name"`
	DistTags    map[string]string `json:"dist-tags"`
	License     string            `json:"license"`
	Author      Author            `json:"author"`
	Homepage    string            `json:"homepage"`
	Description string            `json:"description"`
}

type PackageJson struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func ScanPackageJson(file []byte, conn *websocket.Conn) {
	log.Print(file)

	var npmPackage PackageJson

	json.Unmarshal(file, &npmPackage)

	log.Println(npmPackage)
}
