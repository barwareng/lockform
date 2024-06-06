package models

import "time"

type TrustedContact struct {
	ID        int        `json:"id" gorm:"primaryKey;uniqueIndex"`
	TeamID    string     `json:"teamId" gorm:"index"`
	AddedByID string     `json:"addedById" gorm:"index"`
	Value     string     `json:"value" gorm:"index"`
	Label     string     `json:"label"`
	Domain    string     `json:"domain"`
	Type      string     `json:"type"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
