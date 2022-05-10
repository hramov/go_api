package user

import (
	user_entity "api/src/modules/user/entity"
	user_port "api/src/modules/user/port"
)

type UserService struct {
	Repo user_port.UserRepoPort
}

func (us *UserService) Find() []*user_entity.User {
	return us.Repo.Find()
}

func (us *UserService) FindByEmail(email string) *user_entity.User {
	return us.Repo.FindByEmail(email)
}

func (us *UserService) FindBy(field string, data string) *user_entity.User {
	return us.Repo.FindBy(field, data)
}
