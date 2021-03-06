package db

import (
	"github.com/gocql/gocql"
	"moku-moku/clients/cassandra"
	"moku-moku/domain/access_token"
	"moku-moku/utils/errors"
)

const (
	queryGetByID           = "SELECT access_token, expires, user_id FROM access_tokens WHERE access_token=?;"
	queryUpdateExpiration  = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, expires) VALUES (?, ?, ?);"
)

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
	Create(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func NewRepository() DBRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {

	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer session.Close()

	var at access_token.AccessToken
	if err := session.Query(queryGetByID, id).Scan(&at.AccessToken, &at.TokenExpiration, &at.UserId); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NotFoundError("ID not found")
		}
		return nil, errors.InternalServerError(err.Error())
	}

	return &at, nil
}

// Create Create access token in db
func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
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

// UpdateExpiration Update expiration in db
func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()

	if err := session.Query(queryUpdateExpiration,
		at.TokenExpiration,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
