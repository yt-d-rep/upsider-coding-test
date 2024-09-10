package user

type (
	UserService interface {
		Exists(user *User) (bool, error)
	}
	userService struct {
		uRepo UserRepository
	}
)

func (s *userService) Exists(user *User) (bool, error) {
	user, err := s.uRepo.FindByEmail(user.Email())
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}
