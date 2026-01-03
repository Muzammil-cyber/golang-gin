package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	users map[string]string
}

func NewLoginService() LoginService {
	// In a real application, user data would come from a database
	users := map[string]string{
		"admin": "password",
		"user":  "userpass",
	}
	return &loginService{
		users: users,
	}
}

func (s *loginService) Login(username string, password string) bool {
	if pass, exists := s.users[username]; exists {
		return pass == password
	}
	return false
}
