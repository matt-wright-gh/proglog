package main

import (
	"log"

	"github.com/matt-wright-gh/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
