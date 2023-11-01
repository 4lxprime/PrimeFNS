package utils

import (
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"reflect"
	"strconv"
	"strings"

	responses "github.com/4lxprime/Goldmp/internal/responses/types"
	"github.com/valyala/fasthttp"
)

type Version struct {
	Season int
	Build  float64
	CL     string
	Lobby  string
}

func GetVersion(header *fasthttp.RequestHeader) (*Version, error) {
	version := &Version{
		Season: 0,
		Build:  0.0,
		CL:     "",
		Lobby:  "",
	}

	if header.UserAgent() != nil {
		UA := string(header.UserAgent())
		if !strings.Contains(UA, "-") {
			return nil, fmt.Errorf("bad user-agents")
		}

		parts := strings.Split(UA, "-")
		if len(parts) > 3 {
			buildId := strings.Split(parts[3], ",")[0]
			if _, err := strconv.Atoi(buildId); err == nil {
				version.CL = buildId

			} else {
				buildId = strings.Split(parts[3], " ")[0]
				if _, err := strconv.Atoi(buildId); err == nil {
					version.CL = buildId
				}
			}
		}

		if !strings.Contains(UA, "Release-") {
			return nil, fmt.Errorf("bad user-agents")
		}

		build := strings.Split(strings.Split(UA, "Release-")[1], "-")[0]
		parts = strings.Split(build, ".")
		if len(parts) == 3 {
			build = fmt.Sprintf("%s.%s%s", parts[0], parts[1], parts[2])
		}

		season, err := strconv.Atoi(strings.Split(build, ".")[0])
		if err != nil {
			return nil, fmt.Errorf("error converting season")
		}

		buildFloat, err := strconv.ParseFloat(build, 64)
		if err != nil {
			return nil, fmt.Errorf("error converting build")
		}

		version.Season = season
		version.Build = buildFloat
		version.Lobby = fmt.Sprintf("LobbySeason%d", season)

	} else {
		version.Season = 2
		version.Build = 2.0
		version.Lobby = "LobbyWinterDecor"
	}

	return version, nil
}

func HashString(hasher hash.Hash, data string) string {
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetShop() (*responses.CatalogResponse, error) {
	catalog := &responses.CatalogResponse{}
	catalogFile, err := os.ReadFile("internal/responses/catalog.json")
	if err != nil {
		return nil, err
	}
	catalog.Parse(catalogFile)

	catalogConfig := &responses.CatalogConfigResponse{}
	catalogConfigFile, err := os.ReadFile("internal/responses/catalog_config.json")
	if err != nil {
		return nil, err
	}
	catalogConfig.Parse(catalogConfigFile)

	for _, value := range []responses.CatalogConfigItem{catalogConfig.Daily1, catalogConfig.Daily2, catalogConfig.Daily3, catalogConfig.Daily4, catalogConfig.Daily5, catalogConfig.Daily6, catalogConfig.Featured1, catalogConfig.Featured2} {
		if len(value.ItemGrants) != 0 && len(value.ItemGrants) > 0 {
			var CatalogEntry responses.CatalogEntry
			valueReflect := reflect.ValueOf(value)

			if strings.HasPrefix(strings.ToLower(valueReflect.Type().Name()), "daily") {
				for i, storefront := range catalog.Storefronts {
					if storefront.Name == "BRDailyStorefront" {
						CatalogEntry.Requirements = []responses.CatalogEntryRequirement{}
						CatalogEntry.ItemGrants = []responses.CatalogEntryItemGrant{}

						for x := range value.ItemGrants {
							if len(value.ItemGrants[x]) != 0 && len(value.ItemGrants[x]) > 0 {
								CatalogEntry.DevName = value.ItemGrants[0]
								CatalogEntry.OfferId = value.ItemGrants[0]

								var requirement responses.CatalogEntryRequirement

								requirement.RequirementType = "DenyOnItemOwnership"
								requirement.RequiredId = value.ItemGrants[x]
								requirement.MinQuantity = 1
								CatalogEntry.Requirements = append(CatalogEntry.Requirements, requirement)

								var itemGrant responses.CatalogEntryItemGrant

								itemGrant.TemplateId = value.ItemGrants[x]
								itemGrant.Quantity = 1
								CatalogEntry.ItemGrants = append(CatalogEntry.ItemGrants, itemGrant)
							}
						}

						CatalogEntry.Prices[0].BasePrice = value.Price
						CatalogEntry.Prices[0].RegularPrice = value.Price
						CatalogEntry.Prices[0].FinalPrice = value.Price

						CatalogEntry.SortPriority = -1

						if len(CatalogEntry.ItemGrants) != 0 {
							catalog.Storefronts[i].CatalogEntries = append(catalog.Storefronts[i].CatalogEntries, CatalogEntry)
						}
					}
				}
			}

			if strings.HasPrefix(strings.ToLower(valueReflect.Type().Name()), "featured") {
				for i, storefront := range catalog.Storefronts {
					if storefront.Name == "BRWeeklyStorefront" {
						CatalogEntry.Requirements = []responses.CatalogEntryRequirement{}
						CatalogEntry.ItemGrants = []responses.CatalogEntryItemGrant{}

						for x := range value.ItemGrants {
							if len(value.ItemGrants[x]) != 0 && len(value.ItemGrants[x]) > 0 {
								CatalogEntry.DevName = value.ItemGrants[0]
								CatalogEntry.OfferId = value.ItemGrants[0]

								var requirement responses.CatalogEntryRequirement

								requirement.RequirementType = "DenyOnItemOwnership"
								requirement.RequiredId = value.ItemGrants[x]
								requirement.MinQuantity = 1
								CatalogEntry.Requirements = append(CatalogEntry.Requirements, requirement)

								var itemGrant responses.CatalogEntryItemGrant

								itemGrant.TemplateId = value.ItemGrants[x]
								itemGrant.Quantity = 1
								CatalogEntry.ItemGrants = append(CatalogEntry.ItemGrants, itemGrant)
							}
						}

						CatalogEntry.Prices[0].BasePrice = value.Price
						CatalogEntry.Prices[0].RegularPrice = value.Price
						CatalogEntry.Prices[0].FinalPrice = value.Price

						CatalogEntry.Meta.TileSize = "Normal"
						CatalogEntry.MetaInfo[1].Value = "Normal"

						if len(CatalogEntry.ItemGrants) != 0 {
							catalog.Storefronts[i].CatalogEntries = append(catalog.Storefronts[i].CatalogEntries, CatalogEntry)
						}
					}
				}
			}
		}
	}

	return catalog, nil
}

func GetContentPages(header *fasthttp.RequestHeader) (*responses.ContentpagesResponse, error) {
	version, err := GetVersion(header)
	if err != nil {
		return nil, err
	}

	contentpages := &responses.ContentpagesResponse{}
	contentpagesFile, err := os.ReadFile("internal/responses/contentpages.json")
	if err != nil {
		return nil, err
	}
	contentpages.Parse(contentpagesFile)

	lang := "en"

	if header.Peek("Accept-Language") != nil {
		lg := string(header.Peek("Accept-Language")) // unuseful alloc

		if strings.Contains(lg, "-") && lg != "es-419" && lg != "pt-BR" {
			lang = strings.Split(lg, "-")[0]

		} else {
			lang = lg
		}
	}

	modes := []string{"saveTheWorldUnowned", "battleRoyale", "creative", "saveTheWorld"}
	news := []string{"savetheworldnews", "battleroyalenews"}
	motdnews := []string{"battleroyalenews"}

	contentsCpy := reflect.ValueOf(contentpages)
	contents := contentsCpy.Elem()

	for _, mode := range modes {
		modeField := contents.FieldByName("Subgameselectdata").FieldByName(mode)
		if modeField.IsValid() {
			messageField := modeField.FieldByName("Message")
			if messageField.IsValid() {
				titleField := messageField.FieldByName("Title")
				if titleField.IsValid() && titleField.CanSet() {
					titleField.SetString(titleField.FieldByName(lang).String())
				}

				bodyField := messageField.FieldByName("Body")
				if bodyField.IsValid() && bodyField.CanSet() {
					bodyField.SetString(bodyField.FieldByName(lang).String())
				}
			}
		}
	}

	if version.Build < 5.3 {
		for _, new := range news {
			newField := contents.FieldByName(new)
			if newField.IsValid() {
				messagesField := newField.FieldByName("News").FieldByName("Messages")
				if messagesField.IsValid() {
					iter := messagesField.MapRange()
					for iter.Next() {
						v := iter.Value()
						v.FieldByName("Image").SetString("test")
					}
				}
			}
		}
	}

	for _, motd := range motdnews {
		motdnewField := contents.FieldByName(motd)
		if motdnewField.IsValid() {
			motdField := motdnewField.FieldByName("News").FieldByName("Motds")
			iter := motdField.MapRange()
			for iter.Next() {
				v := iter.Value()
				titleField := v.FieldByName("Title")
				if titleField.IsValid() && titleField.CanSet() {
					titleField.SetString(titleField.FieldByName(lang).String())
				}

				bodyField := v.FieldByName("Body")
				if bodyField.IsValid() && bodyField.CanSet() {
					bodyField.SetString(bodyField.FieldByName(lang).String())
				}
			}
		}
	}

	// todo: set bg

	return contentpages, nil
}
