package accesstoken

import "bookstore_oauth-api/src/utils/errors"

// Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

// Service interface
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService will instantiate a Repository
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
