package db

import (
	"moku-moku/clients/cassandra"
	"moku-moku/domain/access_token"
	"moku-moku/utils/errors"
)

const (
	queryUpdateExpiration = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, token_expiration) VALUES (?, ?, ?);"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpiration(access_token.AccessToken) *errors.RestErr
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

// Update expiration in db
func (r *dbRespository) UpdateExpiration(at access_token.AccessToken) *errors.RestErr {
func (r *dbRespository) Create(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpiration,
		at.AccessToken,
		at.TokenExpiration,
	).Exec(); err != nil {
	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.TokenExpiration).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
