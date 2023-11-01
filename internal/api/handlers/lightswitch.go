package handlers

import "github.com/gofiber/fiber/v2"

type launcherInfoDTO struct {
	AppName       string `json:"appName"`
	CatalogItemID string `json:"catalogItemId"`
	Namespace     string `json:"namespace"`
}

type lightSwitchStatus struct {
	ServiceInstanceID  string          `json:"serviceInstanceId"`
	Status             string          `json:"status"`
	Message            string          `json:"message"`
	MaintenanceURI     interface{}     `json:"maintenanceUri"`
	OverrideCatalogIds []string        `json:"overrideCatalogIds"`
	AllowedActions     []string        `json:"allowedActions"`
	Banned             bool            `json:"banned"`
	LauncherInfoDTO    launcherInfoDTO `json:"launcherInfoDTO"`
}

// ["/lightswitch/api/service/bulk/status"]
func HandleLightswitchServiceBStatus(c *fiber.Ctx) error {
	status := []lightSwitchStatus{
		{
			ServiceInstanceID: "fortnite",
			Status:            "UP",
			Message:           "fortnite is up.",
			MaintenanceURI:    nil,
			OverrideCatalogIds: []string{
				"a7f138b2e51945ffbfdacc1af0541053",
			},
			AllowedActions: []string{
				"PLAY",
				"DOWNLOAD",
			},
			Banned: false,
			LauncherInfoDTO: launcherInfoDTO{
				AppName:       "Fortnite",
				CatalogItemID: "4fe75bbc5a674f4f9b356b5c90567da5",
				Namespace:     "fn",
			},
		},
	}

	return c.JSON(status)
}
