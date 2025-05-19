package validator

import (
	"fmt"
	typesimpo "judo/internal/types"
	"net/url"
)

func ValidUrl(str interface{}) error {
	switch v := str.(type) {
	case typesimpo.LinkCreateRequest:
		if _, err := url.ParseRequestURI(v.URL); err != nil {
			return fmt.Errorf("invalid url: %w", err)
		}

	}
	return nil
}
