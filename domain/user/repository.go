package user

type (
	UserRepository interface {
		FindByEmail(email Email) (*User, error)
		Save(user *User) error
	}
)
