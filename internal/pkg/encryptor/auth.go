package encryptor

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Token user credentials token
type Token struct {
	AccessToken string
	ExpiredAt   int64
}

// Claims user claims
type Claims struct {
	AccountUID  string
	Email       string
	Username    string
	MessengerID string
	Role        int64
	Membership  int64
	jwt.StandardClaims
}

// SignToToken sign claims to token
func (claims *Claims) SignToToken(base64Secret string) (*Token, error) {
	tokenStr, err := genSignedToken(base64Secret, claims)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: tokenStr,
		ExpiredAt:   claims.ExpiresAt,
	}, nil
}

func genSignedToken(base64Secret string, claims jwt.Claims) (string, error) {
	rawSecret, err := Base64Decode(base64Secret)
	if err != nil {
		return "", err
	}

	unSignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := unSignedToken.SignedString(rawSecret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
