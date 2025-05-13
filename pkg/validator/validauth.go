package validator

import (
	"errors"
	"judo/internal/handlers/payload"
	"regexp"
)

func ValidBody(data interface{}) error {
	switch v := data.(type) {
	case payload.LoginRequest:
		if !isValidEmail(v.Email) {
			return errors.New("invalid email")
		}
		if v.Password == "" {
			return errors.New("empty password, required ")
		}
	case payload.RegisterRequest:
		if !isValidEmail(v.Email) {
			return errors.New("invalid email")
		}
		if v.Password == "" {
			return errors.New("empty password, required ")

		}
		if v.Username == "" {
			return errors.New("empty username")
		}
	default:
		return errors.New("invalid type")
	}
	return nil
}

func isValidEmail(email string) bool {
	res := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return res.MatchString(email)
}
