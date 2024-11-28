package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-loc/repository"
)

func NearestLocation(c *fiber.Ctx) error {
	var body RequestBody
	if err := json.Unmarshal(c.Body(), &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Find the nearest road
	road, err := repository.FindNearestRoad(body.Latitude, body.Longitude, body.MaxDistance)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No roads found near the given coordinates",
		})
	}

	// Return the result
	return c.JSON(fiber.Map{
		"name":        road.Properties.Name,
		"coordinates": road.Geometry.Coordinates,
	})
}
