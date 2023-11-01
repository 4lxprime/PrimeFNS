package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	accountId    = "test"
	defaultToken = "primetoken"
)

type oauthV1Response struct {
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	ExpiresAt        time.Time `json:"expires_at"`
	TokenType        string    `json:"token_type"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpires   int       `json:"refresh_expires"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
	AccountID        string    `json:"account_id"`
	ClientID         string    `json:"client_id"`
	InternalClient   bool      `json:"internal_client"`
	ClientService    string    `json:"client_service"`
	DisplayName      string    `json:"displayName"`
	App              string    `json:"app"`
	InAppID          string    `json:"in_app_id"`
	DeviceID         string    `json:"device_id"`
}

type oauthV1Verify struct {
	Token          string    `json:"token"`
	SessionID      string    `json:"session_id"`
	TokenType      string    `json:"token_type"`
	ClientID       string    `json:"client_id"`
	InternalClient bool      `json:"internal_client"`
	ClientService  string    `json:"client_service"`
	AccountID      string    `json:"account_id"`
	ExpiresIn      int       `json:"expires_in"`
	ExpiresAt      time.Time `json:"expires_at"`
	AuthMethod     string    `json:"auth_method"`
	DisplayName    string    `json:"display_name"`
	App            string    `json:"app"`
	InAppID        string    `json:"in_app_id"`
	DeviceID       string    `json:"device_id"`
}

type oauthAccount struct {
	ID                         string `json:"id"`
	DisplayName                string `json:"displayName"`
	Name                       string `json:"name"`
	Email                      string `json:"email"`
	FailedLoginAttempts        int    `json:"failedLoginAttempts"`
	LastLogin                  string `json:"lastLogin"`
	NumberOfDisplayNameChanges int    `json:"numberOfDisplayNameChanges"`
	AgeGroup                   string `json:"ageGroup"`
	Headless                   bool   `json:"headless"`
	Country                    string `json:"country"`
	LastName                   string `json:"lastName"`
	PreferredLanguage          string `json:"preferredLanguage"`
	CanUpdateDisplayName       bool   `json:"canUpdateDisplayName"`
	TfaEnabled                 bool   `json:"tfaEnabled"`
	EmailVerified              bool   `json:"emailVerified"`
	MinorVerified              bool   `json:"minorVerified"`
	MinorExpected              bool   `json:"minorExpected"`
	MinorStatus                string `json:"minorStatus"`
	CabinedMode                bool   `json:"cabinedMode"`
	HasHashedEmail             bool   `json:"hasHashedEmail"`
}

// ["/account/api/oauth/token"]
func HandleOauthToken(c *fiber.Ctx) error {
	// here you can handle the oauth request
	// like switching between method,
	// register user, change accountId global, etc

	response := oauthV1Response{
		AccessToken:      defaultToken,
		ExpiresIn:        28800,
		ExpiresAt:        time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
		TokenType:        "bearer",
		RefreshToken:     defaultToken,
		RefreshExpires:   86400,
		RefreshExpiresAt: time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
		AccountID:        accountId, // accountId
		ClientID:         "",        // clientId
		InternalClient:   true,
		ClientService:    "fortnite",
		DisplayName:      accountId, // accountId
		App:              "fortnite",
		InAppID:          accountId, // accountId
		DeviceID:         "",        // deviceId
	}

	return c.JSON(response)
}

// ["/account/api/oauth/verify"]
func HandleOauthVerify(c *fiber.Ctx) error {
	response := oauthV1Verify{
		Token:          "", // Token
		SessionID:      "", // sessionId
		TokenType:      "bearer",
		ClientID:       "", // clientId
		InternalClient: true,
		ClientService:  "fortnite",
		AccountID:      "", // accountId
		ExpiresIn:      28800,
		ExpiresAt:      time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
		AuthMethod:     "exchange_code",
		DisplayName:    "", // accountId
		App:            "fortnite",
		DeviceID:       "", // deviceId
	}

	return c.JSON(response)
}

// ["/account/api/public/account/:accountId"]
func HandleOauthPublicAccount(c *fiber.Ctx) error {
	//accountId = c.Params("accountId")

	response := oauthAccount{
		ID:                         accountId, // accountId
		DisplayName:                accountId, // accountId
		Name:                       accountId, // accountId
		Email:                      fmt.Sprintf("%s@primebackend.com", accountId),
		FailedLoginAttempts:        0,
		LastLogin:                  time.Now().Format(time.RFC3339),
		NumberOfDisplayNameChanges: 0,
		AgeGroup:                   "UNKNOWN",
		Headless:                   false,
		Country:                    "US",
		LastName:                   "primeBackend",
		PreferredLanguage:          "en",
		CanUpdateDisplayName:       false,
		TfaEnabled:                 false,
		EmailVerified:              true,
		MinorVerified:              false,
		MinorExpected:              false,
		MinorStatus:                "NOT_MINOR",
		CabinedMode:                false,
		HasHashedEmail:             false,
	}

	return c.JSON(response)
}

// ["/account/api/public/account/:accountId/externalAuths"]
func HandleOauthPublicAccountEAuth(c *fiber.Ctx) error {
	return c.JSON([]fiber.Map{})
}

// ["/account/api/oauth/sessions/kill"]
func HandleOauthSessionsKill(c *fiber.Ctx) error {
	return c.SendStatus(204)
}

// ["/account/api/oauth/sessions/kill/:accountId"]
func HandleOauthSessionsKillAccount(c *fiber.Ctx) error {
	return c.SendStatus(204)
}

// ["/fortnite/api/game/v2/tryPlayOnPlatform/account/:accountId"]
func HandleOauthTryPlayOnPlatform(c *fiber.Ctx) error {
	return c.SendString("true")
}
