package db

import (
	"moku-moku/domain/access_token"
	"moku-moku/utils/errors"
)

const (
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, token_expiration) VALUES (?, ?, ?);"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRespository struct {
}

func NewRepository() DBRepository {
	return &dbRespository{}
}

func (r *dbRespository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
