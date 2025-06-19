package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func BuildToken(data map[string]interface{}, privateKey string) (string, int64, error) {
	claims := mapData(data)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", 0, err
	}

	tokenString, err := token.SignedString(pkey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, claims["exp"].(int64), nil
}

func mapData(data map[string]interface{}) jwt.MapClaims {
	mapClaim := jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(3 * time.Hour).Unix(),
		"iss": ISSUER,
		"aud": []string{AUDIENCE},
	}

	for key, value := range data {
		mapClaim[key] = value
	}

	return mapClaim
}
