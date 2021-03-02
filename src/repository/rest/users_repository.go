package rest

import (
	"bookstore_oauth-api/src/domain/users"
	"bookstore_oauth-api/src/utils/errors"
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type userRepository struct{}

func NewRepository() RestUserRepository {
	return &userRepository{}
}

func (ur *userRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	req := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	resp := usersRestClient.Post("/users/login", req)
	if resp == nil || resp.Response == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if resp.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(resp.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(resp.bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error while trying to unmarshal users response")
	}
	return &user, nil
}
