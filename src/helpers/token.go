package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecrets = []byte(Godotenv("JWT_KEYS"))

type claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

func NewToken(email, role string) *claims {
	return &claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(mySecrets)
}

func CheckToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})
	if err != nil {
		return nil, err
	}
	claims := tokens.Claims.(*claims)
	return claims, err
}
