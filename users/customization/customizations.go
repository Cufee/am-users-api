package customization

type ExternalCustomizations map[string]interface{}

type InternalCustomizationOption struct {
	Key             string      `json:"key"`
	Value           interface{} `json:"value"`
	PlusRequired    bool        `json:"plusRequired"`
	PremiumRequired bool        `json:"premiumRequired"`
}
