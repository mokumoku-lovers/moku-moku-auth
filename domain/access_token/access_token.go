package access_token

import "time"

const (
	expirationHours = 24
)

type AccessToken struct {
	AccessToken     string `json:"access_token"`
	UserId          int64  `json:"user_id"`
	TokenExpiration int64  `json:"token_expiration"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		TokenExpiration: time.Now().UTC().Add(expirationHours * time.Hour).Unix(),
	}
}
