package access_token

const (
	expirationHours = 24
)

type AccessToken struct {
	AccessToken     string `json:"access_token"`
	UserId          int64  `json:"user_id"`
	TokenExpiration int64  `json:"token_expiration"`
}
