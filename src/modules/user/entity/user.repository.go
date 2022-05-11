package user_entity

import (
	ioc "api/src/core/container"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func CreateRepository() *UserRepository {
	repository := &UserRepository{
		Db: ioc.Pick[*gorm.DB]("postgres"),
	}
	ioc.Put("UserRepository", repository)
	return repository
}

func (ur *UserRepository) Find() []*User {
	var users []*User
	ur.Db.Find(&users)
	return users
}

func (ur *UserRepository) FindByEmail(email string) (*User, error) {
	var user *User
	result := ur.Db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (ur *UserRepository) FindBy(field string, email string) *User {
	var user *User
	condition := fmt.Sprintf("%s = ?", field)
	ur.Db.Where(condition, email).First(&user)
	return user
}
