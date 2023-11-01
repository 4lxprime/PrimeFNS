package profiles

import (
	"encoding/json"
	"time"
)

type CommonCorePurchase struct {
	PurchaseID         string        `json:"purchaseId"`
	OfferID            string        `json:"offerId"`
	PurchaseDate       time.Time     `json:"purchaseDate"`
	FreeRefundEligible bool          `json:"freeRefundEligible"`
	Fulfillments       []interface{} `json:"fulfillments"`
	LootResult         []struct {
		ItemType    string `json:"itemType"`
		ItemGUID    string `json:"itemGuid"`
		ItemProfile string `json:"itemProfile"`
		Quantity    int    `json:"quantity"`
	} `json:"lootResult"`
	TotalMtxPaid int `json:"totalMtxPaid"`
	Metadata     struct {
		MtxAffiliate   string `json:"mtx_affiliate"`
		MtxAffiliateID string `json:"mtx_affiliate_id"`
	} `json:"metadata"`
	GameContext string `json:"gameContext"`
}

type CommonCoreStats struct {
	Attributes struct {
		SurveyData           struct{} `json:"survey_data"`
		PersonalOffers       struct{} `json:"personal_offers"`
		IntroGamePlayed      bool     `json:"intro_game_played"`
		ImportFriendsClaimed struct{} `json:"import_friends_claimed"`
		MtxPurchaseHistory   struct {
			RefundsUsed   int                  `json:"refundsUsed"`
			RefundCredits int                  `json:"refundCredits"`
			Purchases     []CommonCorePurchase `json:"purchases"`
		} `json:"mtx_purchase_history"`
		UndoCooldowns         []interface{} `json:"undo_cooldowns"`
		MtxAffiliateSetTime   string        `json:"mtx_affiliate_set_time"`
		InventoryLimitBonus   int           `json:"inventory_limit_bonus"`
		CurrentMtxPlatform    string        `json:"current_mtx_platform"`
		MtxAffiliate          string        `json:"mtx_affiliate"`
		ForcedIntroPlayed     string        `json:"forced_intro_played"`
		WeeklyPurchases       struct{}      `json:"weekly_purchases"`
		DailyPurchases        struct{}      `json:"daily_purchases"`
		BanHistory            struct{}      `json:"ban_history"`
		InAppPurchases        struct{}      `json:"in_app_purchases"`
		Permissions           []interface{} `json:"permissions"`
		UndoTimeout           string        `json:"undo_timeout"`
		MonthlyPurchases      struct{}      `json:"monthly_purchases"`
		AllowedToSendGifts    bool          `json:"allowed_to_send_gifts"`
		MfaEnabled            bool          `json:"mfa_enabled"`
		AllowedToReceiveGifts bool          `json:"allowed_to_receive_gifts"`
		GiftHistory           struct{}      `json:"gift_history"`
	} `json:"attributes"`
}

type CommonCoreProfile struct {
	ID              string          `json:"_id"`
	Created         time.Time       `json:"created"`
	Updated         time.Time       `json:"updated"`
	Rvn             int             `json:"rvn"`
	WipeNumber      int             `json:"wipeNumber"`
	AccountID       string          `json:"accountId"`
	ProfileID       string          `json:"profileId"`
	Version         string          `json:"version"`
	Items           []ProfileItem   `json:"items"`
	Stats           CommonCoreStats `json:"stats"`
	CommandRevision int             `json:"commandRevision"`
}

func (commonCore *CommonCoreProfile) Parse(b []byte) error {
	return json.Unmarshal(b, commonCore)
}

// Profile interface -->

func (p CommonCoreProfile) GetId() *string {
	return &p.ID
}

func (p CommonCoreProfile) GetCreated() *time.Time {
	return &p.Created
}

func (p CommonCoreProfile) GetUpdated() *time.Time {
	return &p.Updated
}

func (p CommonCoreProfile) GetRvn() *int {
	return &p.Rvn
}

func (p CommonCoreProfile) GetAccountId() *string {
	return &p.AccountID
}

func (p CommonCoreProfile) GetProfileId() *string {
	return &p.ProfileID
}

func (p CommonCoreProfile) GetItems() interface{} {
	return &p.Items
}

func (p CommonCoreProfile) GetStats() interface{} {
	return p.Stats
}

func (p CommonCoreProfile) GetCommandRevision() *int {
	return &p.CommandRevision
}
