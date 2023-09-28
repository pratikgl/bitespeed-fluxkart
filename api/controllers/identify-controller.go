package controllers

import (
	"fmt"
	"net/http"

	"github.com/pratikgl/bitespeed-fluxkart/api/models"
	"github.com/pratikgl/bitespeed-fluxkart/pkg/utils"
)

// IdentifyRequest is the request body for the Identify endpoint.
type IdentifyRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

// IdentifyResponse is the response body for the Identify endpoint.
type IdentifyResponse struct {
	Contact struct {
		PrimaryContactId    uint     `json:"primaryContactId"`
		Emails              []string `json:"emails"`
		PhoneNumbers        []string `json:"phoneNumbers"`
		SecondaryContactIds []uint   `json:"secondaryContactIds"`
	} `json:"contact"`
}

func Identify(w http.ResponseWriter, r *http.Request) {
	// printing the request body
	fmt.Println("Request Body:", r.Body)
	lola := &models.Contact{}
	utils.ParseBody(r, lola)
	fmt.Println("Lola:", lola, *lola.Email, *lola.PhoneNumber)
}
