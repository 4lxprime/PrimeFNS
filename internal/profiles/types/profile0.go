package profiles

import (
	"encoding/json"
	"time"
)

type Profile0Profile struct {
	ID         string        `json:"_id"`
	Created    time.Time     `json:"created"`
	Updated    time.Time     `json:"updated"`
	Rvn        int           `json:"rvn"`
	WipeNumber int           `json:"wipeNumber"`
	AccountID  string        `json:"accountId"`
	ProfileID  string        `json:"profileId"`
	Version    string        `json:"version"`
	Items      []ProfileItem `json:"items"`
	Stats      struct {
		TemplateID string `json:"templateId"`
		Attributes struct {
			NodeCosts struct {
				T1MainNodepageLayer1 struct {
					TokenHomebasepoints int `json:"Token:homebasepoints"`
				} `json:"t1_main_nodepage_layer1"`
			} `json:"node_costs"`
			MissionAlertRedemptionRecord struct {
				LastClaimTimesMap struct {
					General struct {
						MissionAlertGUIDs []interface{} `json:"missionAlertGUIDs"`
						LastClaimedTimes  []interface{} `json:"lastClaimedTimes"`
					} `json:"General"`
					StormLow struct {
						MissionAlertGUIDs []interface{} `json:"missionAlertGUIDs"`
						LastClaimedTimes  []interface{} `json:"lastClaimedTimes"`
					} `json:"StormLow"`
					Halloween struct {
						MissionAlertGUIDs []interface{} `json:"missionAlertGUIDs"`
						LastClaimedTimes  []interface{} `json:"lastClaimedTimes"`
					} `json:"Halloween"`
					Horde struct {
						MissionAlertGUIDs []interface{} `json:"missionAlertGUIDs"`
						LastClaimedTimes  []interface{} `json:"lastClaimedTimes"`
					} `json:"Horde"`
					Storm struct {
						MissionAlertGUIDs []interface{} `json:"missionAlertGUIDs"`
						LastClaimedTimes  []interface{} `json:"lastClaimedTimes"`
					} `json:"Storm"`
				} `json:"lastClaimTimesMap"`
				OldestClaimIndexForCategory []interface{} `json:"oldestClaimIndexForCategory"`
			} `json:"mission_alert_redemption_record"`
			Twitch         struct{} `json:"twitch"`
			ClientSettings struct {
				PinnedQuestInstances []interface{} `json:"pinnedQuestInstances"`
			} `json:"client_settings"`
			Level         int `json:"level"`
			NamedCounters struct {
				SubGameSelectCountCampaign struct {
					CurrentCount        int    `json:"current_count"`
					LastIncrementedTime string `json:"last_incremented_time"`
				} `json:"SubGameSelectCount_Campaign"`
				SubGameSelectCountAthena struct {
					CurrentCount        int    `json:"current_count"`
					LastIncrementedTime string `json:"last_incremented_time"`
				} `json:"SubGameSelectCount_Athena"`
			} `json:"named_counters"`
			DefaultHeroSquadID string `json:"default_hero_squad_id"`
			CollectionBook     struct {
				Pages                  []interface{} `json:"pages"`
				MaxBookXpLevelAchieved int           `json:"maxBookXpLevelAchieved"`
			} `json:"collection_book"`
			QuestManager struct {
				DailyLoginInterval string `json:"dailyLoginInterval"`
				DailyQuestRerolls  int    `json:"dailyQuestRerolls"`
			} `json:"quest_manager"`
			Bans          struct{} `json:"bans"`
			GameplayStats []struct {
				StatName  string `json:"statName"`
				StatValue int    `json:"statValue"`
			} `json:"gameplay_stats"`
			InventoryLimitBonus int      `json:"inventory_limit_bonus"`
			CurrentMtxPlatform  string   `json:"current_mtx_platform"`
			WeeklyPurchases     struct{} `json:"weekly_purchases"`
			DailyPurchases      struct {
				LastInterval string `json:"lastInterval"`
				PurchaseList struct {
					OneF6B613D4B7BAD47D8A93CAEED2C4996 int `json:"1F6B613D4B7BAD47D8A93CAEED2C4996"`
				} `json:"purchaseList"`
			} `json:"daily_purchases"`
			ModeLoadouts []struct {
				LoadoutName     string   `json:"loadoutName"`
				SelectedGadgets []string `json:"selectedGadgets"`
			} `json:"mode_loadouts"`
			InAppPurchases struct {
				Receipts          []string `json:"receipts"`
				FulfillmentCounts struct {
					ZeroA6CB5B346A149F31A4C3FBDF4BBC198  int `json:"0A6CB5B346A149F31A4C3FBDF4BBC198"`
					DEF6D31D416227E7D73F65B27288ED6F     int `json:"DEF6D31D416227E7D73F65B27288ED6F"`
					Eight2ADCC874CFC2D47927208BAE871CF2B int `json:"82ADCC874CFC2D47927208BAE871CF2B"`
					F0033207441AC38CD704718B91B2C8EF     int `json:"F0033207441AC38CD704718B91B2C8EF"`
				} `json:"fulfillmentCounts"`
			} `json:"in_app_purchases"`
			DailyRewards struct {
				NextDefaultReward   int       `json:"nextDefaultReward"`
				TotalDaysLoggedIn   int       `json:"totalDaysLoggedIn"`
				LastClaimDate       time.Time `json:"lastClaimDate"`
				AdditionalSchedules struct {
					Founderspackdailyrewardtoken struct {
						RewardsClaimed int  `json:"rewardsClaimed"`
						ClaimedToday   bool `json:"claimedToday"`
					} `json:"founderspackdailyrewardtoken"`
				} `json:"additionalSchedules"`
			} `json:"daily_rewards"`
			MonthlyPurchases struct{} `json:"monthly_purchases"`
			Xp               int      `json:"xp"`
			Homebase         struct {
				TownName      string `json:"townName"`
				BannerIconID  string `json:"bannerIconId"`
				BannerColorID string `json:"bannerColorId"`
				FlagPattern   int    `json:"flagPattern"`
				FlagColor     int    `json:"flagColor"`
			} `json:"homebase"`
			PacksGranted int `json:"packs_granted"`
		} `json:"attributes"`
	} `json:"stats"`
	CommandRevision int `json:"commandRevision"`
}

func (profile0 *Profile0Profile) Parse(b []byte) error {
	return json.Unmarshal(b, profile0)
}

// Profile interface -->

func (p Profile0Profile) GetId() *string {
	return &p.ID
}

func (p Profile0Profile) GetCreated() *time.Time {
	return &p.Created
}

func (p Profile0Profile) GetUpdated() *time.Time {
	return &p.Updated
}

func (p Profile0Profile) GetRvn() *int {
	return &p.Rvn
}

func (p Profile0Profile) GetAccountId() *string {
	return &p.AccountID
}

func (p Profile0Profile) GetProfileId() *string {
	return &p.ProfileID
}

func (p Profile0Profile) GetItems() interface{} {
	return &p.Items
}

func (p Profile0Profile) GetStats() interface{} {
	return p.Stats
}

func (p Profile0Profile) GetCommandRevision() *int {
	return &p.CommandRevision
}
