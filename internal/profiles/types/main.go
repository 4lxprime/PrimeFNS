package profiles

import "time"

type Profile interface {
	GetId() *string
	GetCreated() *time.Time
	GetUpdated() *time.Time
	GetRvn() *int
	GetAccountId() *string
	GetProfileId() *string
	GetItems() interface{}
	GetStats() interface{}
	GetCommandRevision() *int
}

type ProfileItem struct {
	TemplateID string   `json:"templateId"`
	Attributes struct{} `json:"attributes"`
	Quantity   int      `json:"quantity"`
}

func MatchProfile(p string) Profile {
	switch p {
	case "athena":
		return AthenaProfile{}

	case "profile0":
		return Profile0Profile{}

	case "common_core":
		return CommonCoreProfile{}

	case "common_public":
		return CommonPublicProfile{}

	default:
		return nil
	}
}
