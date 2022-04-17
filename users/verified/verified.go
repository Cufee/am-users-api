package verified

import (
	"time"
)

type DiscordVerified struct {
	ID                    int       `json:"id" bson:"id"`
	Flags                 int       `json:"flags" bson:"flags"`                                 // User account type - https://discord.com/developers/docs/resources/user#user-object-user-flags
	AccessToken           string    `json:"accessToken" bson:"accessToken"`                     // User's access token
	RefreshToken          string    `json:"refreshToken" bson:"refreshToken"`                   // https://discord.com/developers/docs/topics/oauth2#authorization-code-grant-refresh-token-exchange-example
	AccessTokenExpiration time.Time `json:"accessTokenExpiration" bson:"accessTokenExpiration"` // When the access token expires
}

func (v DiscordVerified) IsVerified() bool {
	return v.AccessToken != "" && v.RefreshToken != "" && v.AccessTokenExpiration.After(time.Now())
}

func (v DiscordVerified) CheckToken() bool {
	logs.Critical("Checking token for user %d, this function should not be called", v.ID)
	return false
}
