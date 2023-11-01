package profiles

import (
	"encoding/json"
	"time"
)

type CommonPublicStats struct {
	Attributes struct {
		BannerColor  string `json:"banner_color"`
		HomebaseName string `json:"homebase_name"`
		BannerIcon   string `json:"banner_icon"`
	} `json:"attributes"`
}

type CommonPublicProfile struct {
	ID              string            `json:"_id"`
	Created         time.Time         `json:"created"`
	Updated         time.Time         `json:"updated"`
	Rvn             int               `json:"rvn"`
	WipeNumber      int               `json:"wipeNumber"`
	AccountID       string            `json:"accountId"`
	ProfileID       string            `json:"profileId"`
	Version         string            `json:"version"`
	Items           []ProfileItem     `json:"items"`
	Stats           CommonPublicStats `json:"stats"`
	CommandRevision int               `json:"commandRevision"`
}

func (commonPublic *CommonPublicProfile) Parse(b []byte) error {
	return json.Unmarshal(b, commonPublic)
}

// Profile interface -->

func (p CommonPublicProfile) GetId() *string {
	return &p.ID
}

func (p CommonPublicProfile) GetCreated() *time.Time {
	return &p.Created
}

func (p CommonPublicProfile) GetUpdated() *time.Time {
	return &p.Updated
}

func (p CommonPublicProfile) GetRvn() *int {
	return &p.Rvn
}

func (p CommonPublicProfile) GetAccountId() *string {
	return &p.AccountID
}

func (p CommonPublicProfile) GetProfileId() *string {
	return &p.ProfileID
}

func (p CommonPublicProfile) GetItems() interface{} {
	return &p.Items
}

func (p CommonPublicProfile) GetStats() interface{} {
	return p.Stats
}

func (p CommonPublicProfile) GetCommandRevision() *int {
	return &p.CommandRevision
}
