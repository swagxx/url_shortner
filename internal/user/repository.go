package user

import (
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

func (ur *UserRepository) Create(user *User) (*User, error) {
	res := ur.DataBase.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (ur *UserRepository) Find(email string) (*User, error) {
	var user User
	res := ur.DataBase.DB.First(&user, "email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
