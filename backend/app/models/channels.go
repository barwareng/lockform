package models

import "time"

type Channel struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	TeamID    string    `json:"teamId"`
	Value     string    `json:"value" gorm:"unique;default:null"`
	Label     string    `json:"label"`
	Type      string    `json:"type"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
