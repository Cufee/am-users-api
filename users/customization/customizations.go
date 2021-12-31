package customization

type ExternalCustomizations map[string]interface{}

type InternalCustomizationOption struct {
	Key             string      `json:"key" bson:"key"`
	Value           interface{} `json:"value" bson:"value"`
	PlusRequired    bool        `json:"plusRequired" bson:"plusRequired"`
	PremiumRequired bool        `json:"premiumRequired" bson:"premiumRequired"`
}
