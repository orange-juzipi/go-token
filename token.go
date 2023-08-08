package go_token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type (
	Token interface {
		JwtSign(id int, expireDuration time.Duration) (string, error)
		JwtParse(tokenString string) (*claims, error)
	}

	token struct {
		Secret string
	}

	claims struct {
		ID int
		jwt.RegisteredClaims
	}
)

func (t *token) JwtSign(id int, expireDuration time.Duration) (string, error) {
	claims := claims{
		id,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(t.Secret))
}

func (t *token) JwtParse(tokenString string) (*claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func New(secret string) Token {
	return &token{Secret: secret}
}
