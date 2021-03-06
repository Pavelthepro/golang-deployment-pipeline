package main

import (
	"io"
	"log"
	"net/http"
)

const version string = "2.1.1"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

func main() {
	log.Printf("Listening on port 8000...")
	http.HandleFunc("/", versionHandler)
	http.ListenAndServe(":8000", nil)
}
