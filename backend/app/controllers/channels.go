package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/lockform/app/models"
	"github.com/lockform/pkg/database"
)

func SaveChannel(c *fiber.Ctx) error {
	channel := &models.Channel{}
	if err := c.BodyParser(channel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	channel.TeamID = c.Cookies("teamId")
	if err := database.DB.Save(&channel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  channel,
	})
}
func GetChannels(c *fiber.Ctx) error {
	var channels []models.Channel
	teamId := c.Cookies("teamId")
	if teamId == "" {
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"data":  nil,
		})
	}
	if err := database.DB.Find(&channels, &models.Channel{TeamID: teamId}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  channels,
	})
}

func SearchChannel(c *fiber.Ctx) error {
	team := &models.Team{}
	type Search struct {
		SearchPhrase string `json:"searchPhrase"`
	}
	search := new(Search)
	channel := &models.Channel{}
	if err := c.QueryParser(search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := database.DB.Find(&channel, &models.Channel{Value: search.SearchPhrase}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	log.Info("ID: ", channel.TeamID)
	if channel.TeamID == "" {
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"data":  nil,
		})
	}
	if err := database.DB.Preload("Channels").Find(&team, &models.Team{ID: channel.TeamID}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  team,
	})

}

func RemoveChannel(c *fiber.Ctx) error {
	channel := &models.Channel{}
	if err := c.QueryParser(channel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	if err := database.DB.Delete(&channel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"data":  nil,
	})
}
