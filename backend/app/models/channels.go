package models

import "time"

type Channel struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	TeamID    string    `json:"teamId"`
	Value     string    `json:"value" gorm:"index"`
	Label     string    `json:"label"`
	Category  string    `json:"category"`
	Type      string    `json:"type"`
	Url       string    `json:"url"`
	IsPublic  bool      `json:"isPublic" gorm:"default:true"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}
