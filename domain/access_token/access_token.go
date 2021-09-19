package access_token

import (
	"moku-moku/utils/errors"
	"strings"
	"time"
)

const (
	expirationHours = 24
)

type AccessToken struct {
	AccessToken     string `json:"access_token"`
	UserId          int64  `json:"user_id"`
	TokenExpiration int64  `json:"token_expiration"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.BadRequest("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.BadRequest("invalid user id")
	}
	if at.TokenExpiration <= 0 {
		return errors.BadRequest("invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		TokenExpiration: time.Now().UTC().Add(expirationHours * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.TokenExpiration, 0).Before(time.Now().UTC())
}
