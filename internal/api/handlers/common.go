package handlers

import (
	"encoding/json"
	"os"
	"strings"

	profiles "github.com/4lxprime/Goldmp/internal/profiles/types"
	responses "github.com/4lxprime/Goldmp/internal/responses/types"
	"github.com/4lxprime/Goldmp/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// ["/clearitemsforshop"]
func HandleClearItemsForShop(c *fiber.Ctx) error {
	statChanged := false

	athena := &profiles.AthenaProfile{}
	athenaFile, err := os.ReadFile("internal/profiles/athena.json") // todo: change this by the user profile
	if err != nil {
		return err
	}
	athena.Parse(athenaFile)

	catalogConfig := &responses.CatalogConfigResponse{}
	catalogConfigFile, err := os.ReadFile("internal/responses/catalogConfig.json")
	if err != nil {
		return err
	}
	catalogConfig.Parse(catalogConfigFile)

	for _, config := range []responses.CatalogConfigItem{catalogConfig.Daily1, catalogConfig.Daily2, catalogConfig.Daily3, catalogConfig.Daily4, catalogConfig.Daily5, catalogConfig.Daily6, catalogConfig.Featured1, catalogConfig.Featured2} {
		if config.ItemGrants != nil {
			for _, grant := range config.ItemGrants {
				if grant != "" {
					for key, item := range athena.Items {
						if strings.EqualFold(grant, item.TemplateID) {
							athena.Items = append(athena.Items[:key], athena.Items[key+1:]...) // delete item from the list
							statChanged = true
						}
					}
				}
			}
		}
	}

	athenaBytes, err := json.Marshal(athena)
	if err != nil {
		return err
	}

	if statChanged {
		athena.Rvn++
		athena.CommandRevision++

		os.WriteFile("internal/api/profiles/athena.json", athenaBytes, 0644)

		return c.SendString("Success")
	}

	return c.SendString("Failed, there are no items to remove")
}

// ["/datarouter/api/v1/public/data"]
func HandleDatarouter(c *fiber.Ctx) error {
	// todo: here get every data
	return c.SendStatus(204)
}

// ["/fortnite/api/game/v2/enabled_features"]
func HandleEnabledFeatures(c *fiber.Ctx) error {
	return c.JSON([]fiber.Map{})
}

// ["/fortnite/api/receipts/v1/account/:accountId/receipts"]
func HandleReceipts(c *fiber.Ctx) error {
	return c.JSON([]fiber.Map{})
}

// ["/fortnite/api/game/v2/chat/:accountId/:room/:profileId/pc"]
func HandleChatPC(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{})
}

// ["/content/api/pages/fortnite-game"]
func HandleContentPage(c *fiber.Ctx) error {
	contentPages, err := utils.GetContentPages(&c.Request().Header)
	if err != nil {
		return err
	}

	return c.JSON(contentPages)
}
