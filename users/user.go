package users

import (
	"aftermath.link/repo/am-users-api/users/bans"
	"aftermath.link/repo/am-users-api/users/customization"
	"aftermath.link/repo/am-users-api/users/paidcontent"
	"aftermath.link/repo/am-users-api/users/players"
	"aftermath.link/repo/am-users-api/users/verified"
	"aftermath.link/repo/logs"
)

// User represents a user account, closely based on the Discord API User object.
// https://discord.com/developers/docs/topics/oauth2
type InternalUser struct {
	ID       string                    `json:"_id" bson:"_id"`                         // user ID
	Locale   string                    `json:"locale" bson:"locale"`                   // User's current language
	Verified *verified.DiscordVerified `json:"discordVerified" bson:"discordVerified"` // Verification status of the user

	Player *players.InternalProfile `json:"player" bson:"player"` // Player profile

	PaidContent          *paidcontent.InternalPaidContent                     `json:"paidContent" bson:"paidContent"`                   // Oaid content
	UniqueCustomizations map[string]customization.InternalCustomizationOption `json:"uniqueCustomizations" bson:"uniqueCustomizations"` // User unique customization options
}

// Check if all required fileds are filled
func (u *InternalUser) IsRecordComplete() bool {
	return u.ID != "" && u.Locale != ""
}

// Converts the internal user to an external user.
func (u *InternalUser) Export() ExternalUser {
	var externalUser ExternalUser
	externalUser.ID = u.ID
	externalUser.Locale = u.Locale
	if u.Verified != nil {
		externalUser.Verified = u.Verified.IsVerified()
	}
	if u.Player != nil {
		externalUser.Player = u.Player.Export()
	}
	if u.PaidContent != nil {
		externalUser.PaidContent = u.PaidContent.Export()
	}

	externalUser.addCustomizations(u)
	return externalUser
}

// Adds customization options to the external user based on the internal user and the customization options.
func (exported *ExternalUser) addCustomizations(u *InternalUser) {
	if u.UniqueCustomizations == nil {
		return
	}
	exported.Customizations = make(map[string]interface{})
	for _, customization := range u.UniqueCustomizations {
		if customization.PlusRequired && exported.PaidContent.IsPlusMember {
			exported.Customizations[customization.Key] = customization.Value
			continue
		}
		if customization.PremiumRequired && exported.PaidContent.IsPremiumMember {
			exported.Customizations[customization.Key] = customization.Value
			continue
		}
		if !customization.PlusRequired && !customization.PremiumRequired {
			exported.Customizations[customization.Key] = customization.Value
		}
	}
}

// Represents a user account which is safe to share with other packages.
type ExternalUser struct {
	ID             string                               `json:"id"` // Unique User ID on Discord
	Verified       bool                                 `json:"verified"`
	Locale         string                               `json:"locale"`
	Player         *players.ExternalProfile             `json:"player"`
	PaidContent    *paidcontent.ExtenalPaidContent      `json:"paidContent"`
	Customizations customization.ExternalCustomizations `json:"customizations"`
	IsBanned       bool                                 `json:"isBanned"`
	BanReason      string                               `json:"banReason,omitempty"`
}

func (u *ExternalUser) AddBanDetails() {
	banned, reason, err := bans.FindActiveBanRecordByDiscordID(u.ID)
	u.BanReason = reason
	u.IsBanned = banned

	if err != nil {
		logs.Warning("Failed to get ban details for user %s: %s", u.ID, err.Error())
	}
}
