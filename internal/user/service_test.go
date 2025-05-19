package user

import (
	typesimpo "judo/internal/types"
	"judo/pkg/jwt"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(user *typesimpo.User) (*typesimpo.User, error) {
	return user, nil
}

func (repo *MockUserRepository) Find(email string) (*typesimpo.User, error) {
	return nil, nil
}

func TestAuthService_Register(t *testing.T) {
	const testEmail = "madk1d302@gmail.com"
	authService := NewAuthService(&MockUserRepository{}, &jwt.JWT{})
	email, err := authService.Register(testEmail, "123456", "Алекс")

	if err != nil {
		t.Fatal(err)
	}
	if email != testEmail {
		t.Errorf("email not equal")
	}
}
