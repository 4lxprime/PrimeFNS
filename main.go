package main

import (
	handlers "github.com/4lxprime/Goldmp/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

const port = ":3551"

func main() {
	r := fiber.New(fiber.Config{})

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	fnApiGroup := r.Group("/fortnite/api")
	fnApiGroup.Get("/v2/versioncheck/:platform", handlers.HandleFortniteVCheck)
	fnApiGroup.Get("/receipts/v1/account/:accountId/receipts", handlers.HandleReceipts)
	fnApiGroup.Get("/storefront/v2/catalog", handlers.HandleStorefrontCatalog)
	fnApiGroup.Get("/calendar/v1/timeline", handlers.HandleTimeline)
	fnApiGroup.Get("/game/v2/enabled_features", handlers.HandleEnabledFeatures)
	fnApiGroup.Post("/game/v2/chat/:accountId/:room/:profileId/pc", handlers.HandleChatPC)
	fnApiGroup.Post("/game/v2/profile/:accountId/client/:mcp", handlers.HandleMcpOthers)
	fnApiGroup.Post("/game/v2/tryPlayOnPlatform/account/:accountId", handlers.HandleOauthTryPlayOnPlatform)
	fnApiGroup.Get("/cloudstorage/system", handlers.HandleCloudstorageSystem)
	fnApiGroup.Get("/cloudstorage/user/:accountId", handlers.HandleCloudstorageUser)
	fnApiGroup.Put("/cloudstorage/user/:accountId/:file", handlers.HandleCloudstorageUserFile)
	fnApiGroup.Get("/cloudstorage/system/:file", handlers.HandleCloudstorageSystemFile)

	authGroup := r.Group("/account/api")
	authGroup.Post("/oauth/token", handlers.HandleOauthToken)
	authGroup.Get("/public/account/:accountId", handlers.HandleOauthPublicAccount)
	authGroup.Get("/public/account/:accountId/externalAuths", handlers.HandleOauthPublicAccountEAuth)
	authGroup.Delete("/oauth/sessions/kill", handlers.HandleOauthSessionsKill)
	authGroup.Delete("/oauth/sessions/kill/:accountId", handlers.HandleOauthSessionsKillAccount)
	authGroup.Get("/oauth/verify", handlers.HandleOauthVerify)

	lightswitchGroup := r.Group("/lightswitch/api/service/")
	lightswitchGroup.Get("/bulk/status", handlers.HandleLightswitchServiceBStatus)

	datarouterGroup := r.Group("/datarouter/api")
	datarouterGroup.Post("/v1/public/data", handlers.HandleDatarouter)

	r.Get("/clearitemsforshop", handlers.HandleClearItemsForShop)
	r.Get("/content/api/pages/fortnite-game", handlers.HandleContentPage)

	r.Listen(port)
}

// ["X"]
func HandleX(c *fiber.Ctx) error {
	return nil
}
