package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Value string `json:"value"`
	jwt.StandardClaims
}

func GenerateJWTToken(value string, JWTSecretKey string, expirationMinutes int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)

	claims := &Claims{
		Value: value,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWTToken(tokenString, JWTSecretKey string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) { return []byte(JWTSecretKey), nil })
	if err != nil {
		return claims, err
	}
	if token == nil || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
