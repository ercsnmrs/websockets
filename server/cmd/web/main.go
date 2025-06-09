package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := routes()
	log.Println("Starting channel listener")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting application on port", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
