package request

import (
	payload2 "judo/internal/handlers/payload"
	"judo/internal/link"
	validator2 "judo/pkg/validator"
)

func isValid[T any](payload T) error {
	switch any(payload).(type) {
	case payload2.LoginRequest:
	case payload2.RegisterRequest:
		if err := validator2.ValidBody(payload); err != nil {
			return err
		}
	case link.LinkCreateRequest:
		if err := validator2.ValidUrl(payload); err != nil {
			return err
		}
	}
	return nil
}
