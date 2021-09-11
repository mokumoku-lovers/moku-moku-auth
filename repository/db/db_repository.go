package db

import "moku-moku/utils/errors"

type DBRepository interface {
	GetByID(string) *errors.RestErr
}

type dbRespository struct {
}

func New() DBRepository {
	return &dbRespository{}
}

func (r *dbRespository) GetByID(id string) *errors.RestErr {
	return nil
}
