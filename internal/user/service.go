package user

import (
	"errors"
)

type AuthService struct {
	userRepo *UserRepository
}

func NewAuthService(repo *UserRepository) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) Register(email, password, name string) (string, error) {
	exist, err := s.userRepo.Find(email)
	if err != nil {
		return "", err
	}
	if exist != nil {
		return "", errors.New(ErrUserExist)
	}
	tempUser := &User{
		Email:    email,
		Password: "",
		Username: name,
	}
	_, err = s.userRepo.Create(tempUser)

	if err != nil {
		return "", err
	}
	return tempUser.Email, nil
}
