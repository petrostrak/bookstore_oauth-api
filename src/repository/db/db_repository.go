package db

import (
	"bookstore_oauth-api/src/domain/accesstoken"
	"bookstore_oauth-api/src/utils/errors"
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
	return nil, errors.NewInternalServerError("db conn not implemented yet")
}
