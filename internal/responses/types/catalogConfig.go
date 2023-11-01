package responses

import "encoding/json"

type CatalogConfigItem struct {
	ItemGrants []string `json:"itemGrants"`
	Price      int      `json:"price"`
}

type CatalogConfigResponse struct {
	Daily1    CatalogConfigItem `json:"daily1"`
	Daily2    CatalogConfigItem `json:"daily2"`
	Daily3    CatalogConfigItem `json:"daily3"`
	Daily4    CatalogConfigItem `json:"daily4"`
	Daily5    CatalogConfigItem `json:"daily5"`
	Daily6    CatalogConfigItem `json:"daily6"`
	Featured1 CatalogConfigItem `json:"featured1"`
	Featured2 CatalogConfigItem `json:"featured2"`
}

func (catalog *CatalogConfigResponse) Parse(b []byte) error {
	return json.Unmarshal(b, catalog)
}
