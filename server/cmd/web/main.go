package main

import (
	"log"
	"net/http"
	"os"
	"websockets/cmd/internal/handlers"
)

func main() {
	mux := routes()
	log.Println("Starting channel listener")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting application on port", port)

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()
	
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
