package auth

type AuthService struct{}

func (as *AuthService) Ping() string {
	return "Pong"
}
