package db

import (
	"moku-moku/domain/access_token"
	"moku-moku/utils/errors"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRespository struct {
}

func New() DBRepository {
	return &dbRespository{}
}

func (r *dbRespository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
