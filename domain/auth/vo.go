package auth

import "upsider-base/shared"

type (
	RawPassword    string
	HashedPassword string
	Token          string
)

func NewRawPassword(raw string) (RawPassword, error) {
	u := RawPassword(raw)
	if err := u.Validate(); err != nil {
		return "", err
	}
	return u, nil
}
func (u RawPassword) Validate() error {
	if len(u) < 8 {
		return &shared.ValidationError{Field: "password", Err: "password must be at least 8 characters"}
	}
	return nil
}
func (u RawPassword) String() string {
	return string(u)
}
func (u RawPassword) IsEmpty() bool {
	return u == ""
}

func NewHashedPassword(hashed string) HashedPassword {
	return HashedPassword(hashed)
}

func (u HashedPassword) String() string {
	return string(u)
}
func (u HashedPassword) IsEmpty() bool {
	return u == ""
}

func NewToken(token string) Token {
	return Token(token)
}
func (u Token) String() string {
	return string(u)
}
