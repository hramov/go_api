package user

import (
	ioc "api/src/core/container"
	user_entity "api/src/modules/user/entity"
	user_port "api/src/modules/user/port"

	"gorm.io/gorm"
)

type UserService struct {
	Repo user_port.UserRepoPort
}

func createService() *UserService {
	db := ioc.Pick[*gorm.DB]("postgres")
	service := &UserService{
		&user_entity.UserRepository{
			Db: db,
		},
	}
	ioc.Put("UserService", service)
	return service
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
