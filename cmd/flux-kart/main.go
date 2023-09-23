package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pratikgl/bitespeed-fluxkart/api/routes"
)

func main() {
	r := mux.NewRouter()
	routes.FluxKartRoutes(r)

	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
