package middleware

import (
	"strings"
	"upsider-base/domain/auth"
	"upsider-base/shared"

	"golang.org/x/crypto/bcrypt"
)

type (
	passwordService struct{}
)

func (s *passwordService) Hash(password string) (auth.HashedPassword, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return auth.NewHashedPassword(string(hashed)), nil
}

func (s *passwordService) NewHashedIfValid(hashed string) (auth.HashedPassword, error) {
	if strings.HasPrefix(hashed, "$2a$") || strings.HasPrefix(hashed, "$2b$") || strings.HasPrefix(hashed, "$2y$") {
		return auth.HashedPassword(hashed), nil
	} else {
		return "", &shared.ValidationError{Field: "password", Err: "invalid hashed password"}
	}
}

func (s *passwordService) Match(password auth.RawPassword, hashed auth.HashedPassword) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed.String()), []byte(password.String())); err != nil {
		return false
	}
	return true
}
