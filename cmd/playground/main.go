package main

import (
	"log"
	"net/http"

	"github.com/Josh-Diamond/apiserver-playground/pkg/server"
)

func main() {
	handler, err := server.BuildAPIHandler()
	if err != nil {
		log.Fatalf("Initialization error: %v", err)
	}

	log.Println("Launching rancher/apiserver target on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
