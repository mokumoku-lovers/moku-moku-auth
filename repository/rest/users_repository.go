package rest

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"moku-moku/domain/users"
	"moku-moku/utils/errors"
)

var (
	usersRestClient = resty.Client{
		HostURL: "http://168.138.215.26:9000",
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func (usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {

	resp, err := usersRestClient.R().
		SetBody(users.UserLoginRequest{
			Email:    email,
			Password: password,
		}).
		Get("/users/login")

	if err != nil || resp.IsError() == true {
		return nil, errors.InternalServerError("invalid response when trying to login user")
	}

	if resp.StatusCode() > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(resp.Body(), &restErr)
		if err != nil {
			return nil, errors.InternalServerError("invalid error interface when trying to login user")
		}

		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(resp.Body(), &user); err != nil {
		return nil, errors.InternalServerError("unmarshall users response error")
	}

	return &user, nil
}
