package v1

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Version(router *mux.Router) {
	url := "/api/v1/version"
	router.HandleFunc("/api/v1/version", func(w http.ResponseWriter, r *http.Request) {
		status, err := w.Write([]byte(VERSION))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("handled %d %s", status, url)
	}).Methods(http.MethodGet)
}

const VERSION = "1.0.0"