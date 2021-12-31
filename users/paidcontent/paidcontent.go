package paidcontent

import "time"

// InternalPlusAccount represents all paid content for a user.
type InternalPaidContent struct {
	InternalPlusAccount    `json:"plusAccount" bson:"plusAccount"`
	InternalPremiumAccount `json:"premiumAccount" bson:"premiumAccount"`
}

// Aftermath Plus account
type InternalPlusAccount struct {
	ActivationDate time.Time `json:"activationDate" bson:"activationDate"`
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationDate"`
}

func (p *InternalPlusAccount) IsExpired() bool {
	return p.ExpirationDate.Before(time.Now())
}

// // Aftermath Premium account
type InternalPremiumAccount struct {
	ActivationDate time.Time `json:"activationDate" bson:"activationDate"`
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationDate"`
}

func (p *InternalPremiumAccount) IsExpired() bool {
	return p.ExpirationDate.Before(time.Now())
}

// Converts the internal premium content to an external premium content.
func (p *InternalPaidContent) Export() *ExtenalPaidContent {
	var externalPaidContent ExtenalPaidContent
	externalPaidContent.IsPremiumMember = !p.InternalPremiumAccount.IsExpired()
	externalPaidContent.IsPlusMember = !p.InternalPlusAccount.IsExpired()
	return &externalPaidContent

}

// ExternalPlusAccount represents all paid content for a user.
// Safe to share with other packages.
type ExtenalPaidContent struct {
	IsPremiumMember bool `json:"isPremiumMember" bson:"isPremiumMember"`
	IsPlusMember    bool `json:"isPlusMember" bson:"isPlusMember"`
}
