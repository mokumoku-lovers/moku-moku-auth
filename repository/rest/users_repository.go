package rest

import (
	"encoding/json"
	"moku-moku/domain/users"
	"moku-moku/utils/errors"

	"github.com/go-resty/resty/v2"
)

var (
	// TODO: Change URL once deployed
	usersRestClient = resty.New().SetHostURL("http://127.0.0.1:9000")
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {

	resp, _ := usersRestClient.R().
		SetBody(users.UserLoginRequest{
			Email:    email,
			Password: password,
		}).
		Post("/users/login")

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
