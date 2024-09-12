package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type PaginateParams struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pagination := PaginateParams{}
		c.QueryParser(&pagination)
		if pagination.Page <= 0 {
			pagination.Page = 1
		}

		switch {
		case pagination.PageSize > 100:
			pagination.PageSize = 100
		case pagination.PageSize <= 0:
			pagination.PageSize = 2
		}
		log.Info("pagination", pagination)
		offset := (pagination.Page - 1) * pagination.PageSize
		return db.Offset(offset).Limit(pagination.PageSize)
	}
}
