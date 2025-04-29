package request

import (
	"judo/pkg/handlerset"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		handlerset.HandlerSet(w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}
	if err := isValid(body); err != nil {
		handlerset.HandlerSet(w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}
	return &body, nil
}
