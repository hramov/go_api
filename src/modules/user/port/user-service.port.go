package user_port

import user_entity "api/src/modules/user/entity"

type UserServicePort interface {
	Find() []*user_entity.User
	FindByEmail(email string) *user_entity.User
	FindBy(field string, data string) *user_entity.User
}
