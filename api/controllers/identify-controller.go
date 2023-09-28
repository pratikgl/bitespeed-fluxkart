package controllers

import (
	"fmt"
	"net/http"

	"github.com/pratikgl/bitespeed-fluxkart/api/models"
	"github.com/pratikgl/bitespeed-fluxkart/pkg/utils"
)

func Identify(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.Contact{}
	utils.ParseBody(r, NewUser)
	fmt.Println(r.Body, NewUser)
}
