package users

import (
	"time"

	"aftermath.link/repo/am-users-api/users/customization"
	"aftermath.link/repo/am-users-api/users/paidcontent"
	"aftermath.link/repo/am-users-api/users/players"
)

// User represents a user account, closely based on the Discord API User object.
// https://discord.com/developers/docs/topics/oauth2
type InternalUser struct {
	ID string `json:"_id,omitempty"` // Internal ID, assigned automatically by the database.

	DiscordID string `json:"discordId"` // Unique User ID on Discord
	Flags     int    `json:"flags"`     // User account type - https://discord.com/developers/docs/resources/user#user-object-user-flags

	Locale                string    `json:"locale"`       // User's current language
	AccessToken           string    `json:"accessToken"`  // User's access token
	RefreshToken          string    `json:"refreshToken"` // https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-refresh-token-exchange-example
	AccessTokenExpiration time.Time `json:"accessTokenExpiration"`

	Player players.InternalProfile `json:"player"` // Player profile

	PaidContent          paidcontent.InternalPaidContent                      `json:"paidContent"`          // Oaid content
	UniqueCustomizations map[string]customization.InternalCustomizationOption `json:"uniqueCustomizations"` // User unique customization options
}

// Converts the internal user to an external user.
func (u *InternalUser) Export() ExternalUser {
	var externalUser ExternalUser
	externalUser.ID = u.DiscordID
	externalUser.Locale = u.Locale
	externalUser.Player = u.Player.Export()
	externalUser.PaidContent = u.PaidContent.Export()

	externalUser.addCustomizations(u)
	return externalUser
}

// Adds customization options to the external user based on the internal user and the customization options.
func (exported *ExternalUser) addCustomizations(u *InternalUser) {
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
	Locale         string                               `json:"locale"`
	Player         players.ExternalProfile              `json:"player"`
	PaidContent    paidcontent.ExtenalPaidContent       `json:"paidContent"`
	Customizations customization.ExternalCustomizations `json:"customizations"`
}
