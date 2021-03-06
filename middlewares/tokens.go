package middlewares

import (
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id int, email string) (string, error) {
	type MyCustomClaims struct {
		jwt.StandardClaims
		UserId int    `json:"user_id"`
		Email  string `json:"email"`
	}

	claims := MyCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: JWT_EXPIRES_AT,
			Issuer:    APPLICATION_NAME,
		},
		UserId: id,
		Email:  email,
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SIGNATURE_KEY))
	return tokenString, err
}

func ValidateToken(tokenString string, keyFunc string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(keyFunc), nil
	})

	return token, err
}
