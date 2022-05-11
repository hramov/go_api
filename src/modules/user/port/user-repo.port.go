package user_port

import user_entity "api/src/modules/user/entity"

type UserRepoPort interface {
	Find() []*user_entity.User
	FindByEmail(email string) (*user_entity.User, error)
	FindBy(field string, data string) *user_entity.User
}
