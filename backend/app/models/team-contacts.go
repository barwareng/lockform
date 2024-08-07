package models

import "time"

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
