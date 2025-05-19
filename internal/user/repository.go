package user

import (
	"errors"
	"gorm.io/gorm"
	typesimpo "judo/internal/types"
	"judo/pkg/db"
)

type UserRepository struct {
	DataBase *db.DB
}

func NewUserRepository(dataBase *db.DB) *UserRepository {
	return &UserRepository{
		dataBase,
	}
}

func (ur *UserRepository) Create(user *typesimpo.User) (*typesimpo.User, error) {
	res := ur.DataBase.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepository) Find(email string) (*typesimpo.User, error) {
	var user typesimpo.User
	err := ur.DataBase.DB.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}
