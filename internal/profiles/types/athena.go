package profiles

import (
	"encoding/json"
	"time"
)

type AthenaGoldmpLoadout struct {
	TemplateID string `json:"templateId"`
	Attributes struct {
		LockerSlotsData struct {
			Slots struct {
				MusicPack struct {
					Items []string `json:"items"`
				} `json:"MusicPack"`
				Character struct {
					Items          []string      `json:"items"`
					ActiveVariants []interface{} `json:"activeVariants"`
				} `json:"Character"`
				Backpack struct {
					Items          []string      `json:"items"`
					ActiveVariants []interface{} `json:"activeVariants"`
				} `json:"Backpack"`
				SkyDiveContrail struct {
					Items          []string      `json:"items"`
					ActiveVariants []interface{} `json:"activeVariants"`
				} `json:"SkyDiveContrail"`
				Dance struct {
					Items []string `json:"items"`
				} `json:"Dance"`
				LoadingScreen struct {
					Items []string `json:"items"`
				} `json:"LoadingScreen"`
				Pickaxe struct {
					Items          []string      `json:"items"`
					ActiveVariants []interface{} `json:"activeVariants"`
				} `json:"Pickaxe"`
				Glider struct {
					Items          []string      `json:"items"`
					ActiveVariants []interface{} `json:"activeVariants"`
				} `json:"Glider"`
				ItemWrap struct {
					Items          []string      `json:"items"`
					ActiveVariants []interface{} `json:"activeVariants"`
				} `json:"ItemWrap"`
			} `json:"slots"`
		} `json:"locker_slots_data"`
		UseCount            int    `json:"use_count"`
		BannerIconTemplate  string `json:"banner_icon_template"`
		BannerColorTemplate string `json:"banner_color_template"`
		LockerName          string `json:"locker_name"`
		ItemSeen            bool   `json:"item_seen"`
		Favorite            bool   `json:"favorite"`
	} `json:"attributes"`
	Quantity int `json:"quantity"`
}

type AthenaStats struct {
	Attributes struct {
		PastSeasons         []interface{} `json:"past_seasons"`
		SeasonMatchBoost    int           `json:"season_match_boost"`
		Loadouts            []string      `json:"loadouts"`
		FavoriteVictorypose string        `json:"favorite_victorypose"`
		MfaRewardClaimed    bool          `json:"mfa_reward_claimed"`
		QuestManager        struct {
			DailyLoginInterval time.Time `json:"dailyLoginInterval"`
			DailyQuestRerolls  int       `json:"dailyQuestRerolls"`
		} `json:"quest_manager"`
		BookLevel               int           `json:"book_level"`
		SeasonNum               int           `json:"season_num"`
		FavoriteConsumableemote string        `json:"favorite_consumableemote"`
		BannerColor             string        `json:"banner_color"`
		FavoriteCallingcard     string        `json:"favorite_callingcard"`
		FavoriteCharacter       string        `json:"favorite_character"`
		FavoriteSpray           []interface{} `json:"favorite_spray"`
		BookXp                  int           `json:"book_xp"`
		Battlestars             int           `json:"battlestars"`
		BattlestarsSeasonTotal  int           `json:"battlestars_season_total"`
		StylePoints             int           `json:"style_points"`
		AlienStylePoints        int           `json:"alien_style_points"`
		PartyAssistQuest        string        `json:"party_assist_quest"`
		PinnedQuest             string        `json:"pinned_quest"`
		PurchasedBpOffers       []interface{} `json:"purchased_bp_offers"`
		FavoriteLoadingscreen   string        `json:"favorite_loadingscreen"`
		BookPurchased           bool          `json:"book_purchased"`
		LifetimeWins            int           `json:"lifetime_wins"`
		FavoriteHat             string        `json:"favorite_hat"`
		Level                   int           `json:"level"`
		FavoriteBattlebus       string        `json:"favorite_battlebus"`
		FavoriteMapmarker       string        `json:"favorite_mapmarker"`
		FavoriteVehicledeco     string        `json:"favorite_vehicledeco"`
		AccountLevel            int           `json:"accountLevel"`
		FavoriteBackpack        string        `json:"favorite_backpack"`
		FavoriteDance           []string      `json:"favorite_dance"`
		InventoryLimitBonus     int           `json:"inventory_limit_bonus"`
		LastAppliedLoadout      string        `json:"last_applied_loadout"`
		FavoriteSkydivecontrail string        `json:"favorite_skydivecontrail"`
		FavoritePickaxe         string        `json:"favorite_pickaxe"`
		FavoriteGlider          string        `json:"favorite_glider"`
		DailyRewards            struct{}      `json:"daily_rewards"`
		Xp                      int           `json:"xp"`
		SeasonFriendMatchBoost  int           `json:"season_friend_match_boost"`
		ActiveLoadoutIndex      int           `json:"active_loadout_index"`
		FavoriteMusicpack       string        `json:"favorite_musicpack"`
		BannerIcon              string        `json:"banner_icon"`
		FavoriteItemwraps       []string      `json:"favorite_itemwraps"`
	} `json:"attributes"`
}

type AthenaProfile struct {
	ID              string        `json:"_id"`
	Created         time.Time     `json:"created"`
	Updated         time.Time     `json:"updated"`
	Rvn             int           `json:"rvn"`
	WipeNumber      int           `json:"wipeNumber"`
	AccountID       string        `json:"accountId"`
	ProfileID       string        `json:"profileId"`
	Version         string        `json:"version"`
	Items           []ProfileItem `json:"items"`
	Stats           AthenaStats   `json:"stats"`
	CommandRevision int           `json:"commandRevision"`
}

func (athena *AthenaProfile) Parse(b []byte) error {
	return json.Unmarshal(b, athena)
}

// Profile interface -->

func (p AthenaProfile) GetId() *string {
	return &p.ID
}

func (p AthenaProfile) GetCreated() *time.Time {
	return &p.Created
}

func (p AthenaProfile) GetUpdated() *time.Time {
	return &p.Updated
}

func (p AthenaProfile) GetRvn() *int {
	return &p.Rvn
}

func (p AthenaProfile) GetAccountId() *string {
	return &p.AccountID
}

func (p AthenaProfile) GetProfileId() *string {
	return &p.ProfileID
}

func (p AthenaProfile) GetItems() interface{} {
	return &p.Items
}

func (p AthenaProfile) GetStats() interface{} {
	return p.Stats
}

func (p AthenaProfile) GetCommandRevision() *int {
	return &p.CommandRevision
}
