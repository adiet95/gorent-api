package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecrets = []byte(os.Getenv("JWT_KEYS"))

type Claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

func NewToken(email, role string) *Claims {
	return &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func (c *Claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(mySecrets)
}

func CheckToken(token string) (*Claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})
	if err != nil {
		return nil, err
	}
	Claims := tokens.Claims.(*Claims)
	return Claims, err
}
