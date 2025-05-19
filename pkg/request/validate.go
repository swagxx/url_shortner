package request

import (
	typesimpo "judo/internal/types"
	"judo/pkg/dto"
	"judo/pkg/validator"
)

func isValid[T any](payload T) error {
	switch any(payload).(type) {
	case dto.LoginRequest:
	case dto.RegisterRequest:
		if err := validator.ValidBody(payload); err != nil {
			return err
		}
	case typesimpo.LinkCreateRequest:
		if err := validator.ValidUrl(payload); err != nil {
			return err
		}
	}
	return nil
}
