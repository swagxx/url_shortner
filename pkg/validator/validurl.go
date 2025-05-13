package validator

import (
	"fmt"
	"judo/internal/link"
	"net/url"
)

func ValidUrl(str interface{}) error {
	switch v := str.(type) {
	case link.LinkCreateRequest:
		if _, err := url.ParseRequestURI(v.URL); err != nil {
			return fmt.Errorf("invalid url: %w", err)
		}

	}
	return nil
}
