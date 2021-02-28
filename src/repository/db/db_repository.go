package db

import (
	"bookstore_oauth-api/src/clients/cassandra"
	"bookstore_oauth-api/src/domain/accesstoken"
	"bookstore_oauth-api/src/utils/errors"

	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
)

// NewRepository will return a pointer to a dbRepository struct
func NewRepository() Repository {
	return &dbRepository{}
}

// Repository interface of DB
type Repository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (dr *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result accesstoken.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserID, &result.ClientID, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found for given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}
