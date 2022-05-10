package auth_port

import auth_dto "api/src/modules/auth/dto"

type AuthServicePort interface {
	Login(dto auth_dto.LoginDto) (string, error)
}
