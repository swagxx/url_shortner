package request

import (
	"judo/internal/validator"
)

func isValid[T any](payload T) error {
	if err := validator.ValidBody(payload); err != nil {
		return err
	}
	return nil
}
