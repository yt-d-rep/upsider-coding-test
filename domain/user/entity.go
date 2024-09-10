package user

import (
	"upsider-coding-test/domain/auth"
	"upsider-coding-test/domain/company"
)

type (
	User struct {
		id             UserID
		username       Username
		email          Email
		hashedPassword auth.HashedPassword
		companyID      company.CompanyID
	}
)

func NewUser(username string, email string, hashedPassword auth.HashedPassword, companyID company.CompanyID) (*User, error) {
	id := NewUserID()
	usernameConverted, err := NewUsername(username)
	if err != nil {
		return nil, err
	}
	emailConverted, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	return &User{
		id:             id,
		username:       usernameConverted,
		email:          emailConverted,
		hashedPassword: hashedPassword,
		companyID:      companyID,
	}, nil
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Username() Username {
	return u.username
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) HashedPassword() auth.HashedPassword {
	return u.hashedPassword
}

func (u *User) CompanyID() company.CompanyID {
	return u.companyID
}
