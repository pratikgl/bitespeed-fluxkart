package controllers

import (
	"encoding/json"
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
	// Parse the request body
	var req IdentifyRequest
	utils.ParseRequestBody(r, &req)

	// Getting all the contacts where the email or phone number matches
	if contacts, err := models.GetContactsByEmailOrPhone(req.Email, req.PhoneNumber); err == nil {
		// If there are no contacts, create a new contact
		if len(contacts) == 0 {
			newContact := models.Contact{
				PhoneNumber:    &req.PhoneNumber,
				Email:          &req.Email,
				LinkPrecedence: "primary",
			}
			err := newContact.CreateContact()
			if err != nil {
				fmt.Println(err) // TODO: Throw a useful error on the client side
			}
		} else {
			fmt.Println("multiple contacts are there", contacts)
		}
	} else {
		fmt.Println(err) // TODO: Throw a useful error on the client side
	}

	// returning the response
	contacts, _ := models.GetContactsByEmailOrPhone(req.Email, req.PhoneNumber)
	var res IdentifyResponse
	res.Contact.PrimaryContactId = contacts[0].Id
	for _, contact := range contacts {
		res.Contact.Emails = append(res.Contact.Emails, *contact.Email)
		res.Contact.PhoneNumbers = append(res.Contact.PhoneNumbers, *contact.PhoneNumber)
		if contact.LinkPrecedence == "secondary" {
			res.Contact.SecondaryContactIds = append(res.Contact.SecondaryContactIds, contact.Id)
		}
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	if response, err := json.Marshal(res); err == nil {
		w.Write(response)
	} else {
		fmt.Println(err) // TODO: Throw a useful error on the client side
	}
}
