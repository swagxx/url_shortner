package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
}

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{secret}
}

func (j *JWT) GenerateToken(data JWTData) (string, error) {
	if data.Email == "" {
		return "", fmt.Errorf("email is required")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})
	s, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) ParseToken(tokenString string) (*JWTData, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})

	if err != nil || !token.Valid {
		return nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}
	emailRaw, exists := claims["email"]
	if !exists {
		return nil, false
	}
	email, ok := emailRaw.(string)

	if !ok {
		return nil, false
	}
	return &JWTData{email}, true

}
