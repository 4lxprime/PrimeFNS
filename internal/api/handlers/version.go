package handlers

import "github.com/gofiber/fiber/v2"

// ["/fortnite/api/v2/versioncheck/:platform"]
func HandleFortniteVCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"type": "NO_UPDATE",
	})
}
