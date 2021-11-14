package access_token

import (
	"moku-moku/repository/rest"
	"moku-moku/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	usersRepository rest.RestUsersRepository
	repository      Repository
}

func NewService(usersRepository rest.RestUsersRepository, repository Repository) Service {
	return &service{
		usersRepository: usersRepository,
		repository:      repository,
	}
}

func (s *service) GetByID(accessTokenId string) (*AccessToken, *errors.RestErr) {
	at, err := s.repository.GetByID(accessTokenId)
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *errors.RestErr) {
	// Validation of the login data
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Authenticate the user with User API
	user, err := s.usersRepository.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// New AT generation
	at := GetNewAccessToken(user.Id)
	at.Generate()

	// Store the new AT in Cassandra
	if err := s.repository.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
