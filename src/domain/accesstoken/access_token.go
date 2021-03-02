package accesstoken

import (
	"bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

// AccessTokenRequest struct
type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate validates the access token
func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		break
	case grantTypeClientCredentials:
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}
	// TODO: validate parametes for each grant_type
	return nil
}

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Validate validates the access token
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

// GetNewAccessToken returns the expiration time of the token
func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired checks if the Access Token is expired
func (at *AccessToken) IsExpired() bool {
	// now := time.Now().UTC()
	// expirationTime := time.Unix(at.Expires, 0)
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
