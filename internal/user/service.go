package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	typesimpo "judo/internal/types"
	"judo/pkg/di"
	"judo/pkg/jwt"
)

type AuthService struct {
	userRepo di.IUserRepository
	JWT      *jwt.JWT
}

func NewAuthService(repo di.IUserRepository, j *jwt.JWT) *AuthService {
	return &AuthService{
		userRepo: repo,
		JWT:      j,
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
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	tempUser := &typesimpo.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	_, err = s.userRepo.Create(tempUser)

	token, err := s.JWT.GenerateToken(jwt.JWTData{})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	exits, err := s.userRepo.Find(email)
	if err != nil {
		return "", err
	}
	if exits == nil {
		return "", errors.New(ErrUserNotFound)
	}
	err = bcrypt.CompareHashAndPassword([]byte(exits.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrPasswordNotCorrect)
	}
	token, err := s.JWT.GenerateToken(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}
