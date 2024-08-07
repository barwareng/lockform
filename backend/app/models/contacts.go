package models

import "time"

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
