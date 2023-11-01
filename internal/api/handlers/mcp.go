package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	profiles "github.com/4lxprime/Goldmp/internal/profiles/types"
	"github.com/gofiber/fiber/v2"
)

type mcpResponse struct {
	ProfileRevision            int              `json:"profileRevision"`
	ProfileID                  string           `json:"profileId"`
	ProfileChangesBaseRevision int              `json:"profileChangesBaseRevision"`
	ProfileChanges             []profileChanges `json:"profileChanges"`
	ProfileCommandRevision     int              `json:"profileCommandRevision"`
	ServerTime                 string           `json:"serverTime"`
	ResponseVersion            int              `json:"responseVersion"`
}

type profileChanges struct {
	ChangeType string `json:"changeType"`
	Profile    string `json:"profile"`
}

// ["/fortnite/api/game/v2/profile/:accountId/client/:mcp"]
func HandleMcpOthers(c *fiber.Ctx) error {
	profileId := string(c.Query("profileId")) // unuseful allocation

	profileData, err := os.ReadFile(fmt.Sprintf("internal/profiles/%s.json", profileId))
	if err != nil {
		return err
	}

	profile := profiles.MatchProfile(profileId)
	json.Unmarshal(profileData, profile)

	applyProfileChange := []profileChanges{}
	baseRevision := *profile.GetRvn()
	queryRevision, err := strconv.Atoi(string(c.Query("rvn"))) // unuseful allocation
	if err != nil {
		return err
	}

	if queryRevision != baseRevision {
		applyProfileChange = append(applyProfileChange, profileChanges{
			ChangeType: "fullProfileUpdate",
			Profile:    profileId,
		})
	}

	response := mcpResponse{
		ProfileRevision:            baseRevision,
		ProfileID:                  profileId,
		ProfileChangesBaseRevision: baseRevision,
		ProfileChanges:             applyProfileChange,
		ProfileCommandRevision:     *profile.GetCommandRevision(),
		ServerTime:                 time.Now().Format(time.RFC3339),
		ResponseVersion:            1,
	}

	return c.JSON(response)
}

// ["/fortnite/api/game/v2/profile/:accountId/client/ClientQuestLogin"]
func HandleMcpClientQuestLogin(c *fiber.Ctx) error {
	response := mcpResponse{
		ProfileRevision:            0,
		ProfileID:                  string(c.Query("profileId")), // unuseful alloc
		ProfileChangesBaseRevision: 0,
		ProfileChanges:             []profileChanges{},
		ProfileCommandRevision:     0,
		ServerTime:                 time.Now().Format(time.RFC3339),
		ResponseVersion:            1,
	}

	// todo: handle this

	return c.JSON(response)
}
