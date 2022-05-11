package auth_dto

type LoginDto struct {
	Email    string `json:"email" validate:"type=email"`
	Password string `json:"password" validate:"required,type=password"`
}
