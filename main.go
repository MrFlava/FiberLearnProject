package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Recieved request for items")
}

func main() {
	http.HandleFunc("/items", itemsHandler)

	log.Info("Starting server on port 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
