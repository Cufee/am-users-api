package bans

import (
	"time"

	databasebanrecords "aftermath.link/repo/am-users-api/database/logic/banrecords"
	"aftermath.link/repo/logs"
)

type InternalBan struct {
	ID        int                    `json:"id" bson:"_id,omitempty"`
	UserID    int                    `json:"userId" bson:"userId"`
	Reason    string                 `json:"reason" bson:"reason"`
	ExpiresAt time.Time              `json:"expiration" bson:"expiration"`
	Meta      map[string]interface{} `json:"meta" bson:"meta"`
}

func (b InternalBan) IsExpired() bool {
	return b.ExpiresAt.Before(time.Now())
}

func FindActiveBanRecordByDiscordID(id string) (banned bool, reason string, err error) {
	filter := make(map[string]interface{})
	filter["discordId"] = id
	filter["expiration"] = map[string]interface{}{"$gte": time.Now()}

	sort := make(map[string]interface{})
	sort["expiration"] = -1

	var banRecords []InternalBan
	err = databasebanrecords.FindBanRecordWithFilter(filter, sort, 1, &banRecords)
	if err != nil {
		logs.Warning("databasebanrecords.FindBanRecordWithFilter failed: %s", err)
		return false, "", logs.Wrap(err, "FindBanRecordWithFilter failed")
	}

	if len(banRecords) == 0 {
		return false, "", nil
	}

	return true, banRecords[0].Reason, nil
}
