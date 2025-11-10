package main

import (
	"log"

	"workshop4/internal/app"
)

func main() {
	srv, err := app.New()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	log.Printf("starting server on %s", srv.Addr)
	log.Fatal(srv.Listen())
}
