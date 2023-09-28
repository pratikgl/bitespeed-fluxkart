package models

import (
	"time"

	"github.com/pratikgl/bitespeed-fluxkart/pkg/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

// The json tag in a Go struct is used to specify how a struct
// field should be marshaled to JSON when encoding the struct
// into a JSON representation or unmarshaled from JSON when
// decoding JSON into a Go struct.
type Contact struct {
	Id             uint           `json:"id" gorm:"primaryKey"`
	PhoneNumber    *string        `json:"phoneNumber" gorm:"default:null"`
	Email          *string        `json:"email" gorm:"default:null"`
	LinkedId       *uint          `json:"linkedId" gorm:"default:null"`                                     // This is a foreign key to another Contact
	LinkPrecedence string         `json:"linkPrecedence" gorm:"type:enum('primary', 'secondary');not null"` // either "primary" or "secondary"
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt"`
}

func init() {
	DB = config.Connect()
	DB.AutoMigrate(&Contact{})
}

// CreateContact creates a new contact in the database.
func (contact *Contact) CreateContact() error {
	return DB.Create(contact).Error
}

// Find contacts
func GetContactsByEmailOrPhone(email string, phone string) ([]Contact, error) {
	var contacts []Contact
	query := DB.Where("email = ? OR phone_number = ?", email, phone)

	if err := query.Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

// UpdateContact updates a contact in the database.
func UpdateContact(contact *Contact) error {
	return DB.Save(contact).Error
}
