package controllers

import (
	"net/http"
)

func Identify(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Identify"))
}
