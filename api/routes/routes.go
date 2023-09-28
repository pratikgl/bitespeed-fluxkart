package routes

import (
	"github.com/gorilla/mux"
	"github.com/pratikgl/bitespeed-fluxkart/api/controllers"
)

var FluxKartRoutes = func(router *mux.Router) {
	router.HandleFunc("/identify", controllers.Identify).Methods("POST")
}
