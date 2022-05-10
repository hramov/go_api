package user_entity

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (ur *UserRepository) Find() []*User {
	var users []*User
	ur.Db.Find(&users)
	return users
}

func (ur *UserRepository) FindByEmail(email string) *User {
	var user *User
	ur.Db.Where("email = ?", email).First(&user)
	return user
}

func (ur *UserRepository) FindBy(field string, email string) *User {
	var user *User
	condition := fmt.Sprintf("%s = ?", field)
	ur.Db.Where(condition, email).First(&user)
	return user
}
