package responses

import (
	"encoding/json"
	"time"
)

type CatalogEntryPrice struct {
	CurrencyType    string    `json:"currencyType"`
	CurrencySubType string    `json:"currencySubType"`
	RegularPrice    int       `json:"regularPrice"`
	FinalPrice      int       `json:"finalPrice"`
	SaleExpiration  time.Time `json:"saleExpiration"`
	BasePrice       int       `json:"basePrice"`
}

type CatalogEntryRequirement struct {
	RequirementType string `json:"requirementType"`
	RequiredId      string `json:"requiredId"`
	MinQuantity     int    `json:"minQuantity"`
}

type CatalogEntryItemGrant struct {
	TemplateId string `json:"templateId"`
	Quantity   int    `json:"quantity"`
}

type CatalogEntryMeta struct {
	MaxConcurrentPurchases string `json:"MaxConcurrentPurchases"`
	Preroll                string `json:"Preroll"`
	ProfileID              string `json:"ProfileId"`
	EventLimit             string `json:"EventLimit"`
	TileSize               string `json:"TileSize"`
}

type CatalogEntryMetaInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CatalogEntry struct {
	DevName              string                    `json:"devName"`
	OfferId              string                    `json:"offerId"`
	FulfillmentIds       []interface{}             `json:"fulfillmentIds"`
	DailyLimit           int                       `json:"dailyLimit"`
	WeeklyLimit          int                       `json:"weeklyLimit"`
	MonthlyLimit         int                       `json:"monthlyLimit"`
	Categories           []string                  `json:"categories"`
	Prices               []CatalogEntryPrice       `json:"prices"`
	Meta                 CatalogEntryMeta          `json:"meta"`
	MatchFilter          string                    `json:"matchFilter"`
	FilterWeight         int                       `json:"filterWeight"`
	AppStoreId           []interface{}             `json:"appStoreId"`
	Requirements         []CatalogEntryRequirement `json:"requirements"`
	OfferType            string                    `json:"offerType"`
	GiftInfo             struct{}                  `json:"giftInfo"`
	Refundable           bool                      `json:"refundable"`
	MetaInfo             []CatalogEntryMetaInfo    `json:"metaInfo"`
	DisplayAssetPath     string                    `json:"displayAssetPath"`
	ItemGrants           []CatalogEntryItemGrant   `json:"itemGrants"`
	SortPriority         int                       `json:"sortPriority"`
	CatalogGroupPriority int                       `json:"catalogGroupPriority"`
}

type CatalogStorefront struct {
	Name           string         `json:"name"`
	CatalogEntries []CatalogEntry `json:"catalogEntries"`
}

type CatalogResponse struct {
	RefreshIntervalHrs int                 `json:"refreshIntervalHrs"`
	DailyPurchaseHrs   int                 `json:"dailyPurchaseHrs"`
	Expiration         time.Time           `json:"expiration"`
	Storefronts        []CatalogStorefront `json:"storefronts"`
}

func (catalog *CatalogResponse) Parse(b []byte) error {
	return json.Unmarshal(b, catalog)
}
