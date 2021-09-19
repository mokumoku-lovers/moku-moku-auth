package db

import (
	"moku-moku/clients/cassandra"
	"moku-moku/domain/access_token"
	"moku-moku/utils/errors"
)

const (
	queryUpdateExpiration = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
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

// Update expiration in db
func (r *dbRespository) UpdateExpiration(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpiration,
		at.AccessToken,
		at.TokenExpiration,
	).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
