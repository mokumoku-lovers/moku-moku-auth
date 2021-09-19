package db

import (
	"moku-moku/clients/cassandra"
	"moku-moku/domain/access_token"
	"moku-moku/utils/errors"
)

const (
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, token_expiration) VALUES (?, ?, ?);"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
}

type dbRespository struct {
}

func NewRepository() DBRepository {
	return &dbRespository{}
}

func (r *dbRespository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}

func (r *dbRespository) Create(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.TokenExpiration).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
