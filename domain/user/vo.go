package user

import (
	"regexp"

	"github.com/google/uuid"

	"upsider-base/shared"
)

type (
	UserID   uuid.UUID
	Username string
	Email    string
)

func NewUserID() UserID {
	return UserID(uuid.New())
}
func (u UserID) String() string {
	return uuid.UUID(u).String()
}

func NewUsername(username string) (Username, error) {
	if err := Username(username).Validate(); err != nil {
		return "", err
	}
	return Username(username), nil
}
func (u Username) String() string {
	return string(u)
}
func (u Username) Validate() error {
	if len(u) < 3 {
		return &shared.ValidationError{Field: "username", Err: "username must have at least 3 characters"}
	}
	return nil
}

func NewEmail(email string) (Email, error) {
	if err := Email(email).Validate(); err != nil {
		return "", err
	}
	return Email(email), nil
}
func (u Email) Validate() error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(string(u)) {
		return &shared.ValidationError{Field: "email", Err: "invalid email format"}
	}
	return nil
}
func (u Email) String() string {
	return string(u)
}
