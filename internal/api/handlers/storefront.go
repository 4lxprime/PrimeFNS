package handlers

import (
	"github.com/4lxprime/Goldmp/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// ["/fortnite/api/storefront/v2/catalog"]
func HandleStorefrontCatalog(c *fiber.Ctx) error {
	catalog, err := utils.GetShop()
	if err != nil {
		return err
	}

	return c.JSON(catalog)
}
