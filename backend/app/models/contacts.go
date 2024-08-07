package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Contact struct {
	ID        int        `json:"id" gorm:"primaryKey;uniqueIndex"`
	Value     string     `json:"value" gorm:"index"`
	Label     string     `json:"label"`
	Domain    string     `json:"domain"`
	Url       string     `json:"url"`
	Type      string     `json:"type"`
	Category  string     `json:"category"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
type TeamContact struct {
	ID                  int        `json:"id" gorm:"primaryKey;uniqueIndex"`
	ContactID           int        `json:"contactId" gorm:"index"`
	TeamID              string     `json:"teamId" gorm:"index"`
	AddedByID           string     `json:"addedById" gorm:"index"`
	IsTrusted           bool       `json:"isTrusted" gorm:"index"`
	ReasonForUntrusting string     `json:"reasonForUntrusting"`
	CreatedAt           time.Time  `json:"createdAt"`
	UpdatedAt           time.Time  `json:"updatedAt"`
	DeletedAt           *time.Time `sql:"index" json:"-"`
}

// Response model for fetching contacts
type ContactList struct {
	ID                  int     `json:"id"`
	Value               string  `json:"value"`
	Label               string  `json:"label"`
	Domain              string  `json:"domain"`
	Url                 string  `json:"url"`
	Type                string  `json:"type"`
	Category            string  `json:"category"`
	IsTrusted           bool    `json:"isTrusted"`
	ReasonForUntrusting string  `json:"reasonForUntrusting"`
	AddedBy             AddedBy `json:"addedBy" gorm:"type:added_by"` // Specify custom type if needed
}

type AddedBy struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a AddedBy) Value() (driver.Value, error) {
	return fmt.Sprintf("%s,%s,%s", a.ID, a.Name, a.Email), nil
}

// Scan implements the sql.Scanner interface
func (a *AddedBy) Scan(value interface{}) error {
	if value == nil {
		*a = AddedBy{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("type assertion to string failed")
	}

	parts := strings.SplitN(str, ",", 3)
	if len(parts) != 3 {
		return errors.New("failed to parse AddedBy value")
	}

	a.ID = parts[0]
	a.Name = parts[1]
	a.Email = parts[2]
	return nil
}
