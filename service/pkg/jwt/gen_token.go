package jwtapp

import (
	constant "app/internal/constants"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenToken(data map[string]interface{}, tokenTime constant.TOKEN_TIME) (*string, error) {
	privateKeyData, err := os.ReadFile("key/jwt/private.pem")
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(tokenTime)).Unix(),
	}
	for k, v := range data {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}
