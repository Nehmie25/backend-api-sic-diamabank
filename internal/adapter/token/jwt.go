package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
	expire time.Duration
}

func NewJWTService(secret string, expire time.Duration) *JWTService {
	return &JWTService{
		secret: []byte(secret),
		expire: expire,
	}
}

func (j *JWTService) Generate(userID string, role string, email string, nom string) (string, error) {
	fmt.Println("nom:", nom)
	claims := jwt.MapClaims{
		"id":  userID,
		"role": role,
		"email": email,
		"nom": nom,
		"exp":  time.Now().Add(j.expire).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWTService) Validate(tokenStr string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
