package usecase

import (
	"upsider-coding-test/domain/auth"
	"upsider-coding-test/domain/company"
	"upsider-coding-test/domain/user"
	"upsider-coding-test/shared"
)

type (
	UserUsecase interface {
		Register(input *RegisterInput) (*user.User, error)
		Login(input *LoginInput) (auth.Token, error)
	}
	userUsecase struct {
		userRepository  user.UserRepository
		userService     user.UserService
		passwordService auth.PasswordService
		tokenService    auth.TokenService
	}
	RegisterInput struct {
		Username    string
		Email       string
		RawPassword string
		CompanyID   string
	}
	LoginInput struct {
		Email       string
		RawPassword string
	}
)

func (u *userUsecase) Register(input *RegisterInput) (*user.User, error) {
	hashedPassword, err := u.passwordService.Hash(input.RawPassword)
	if err != nil {
		return nil, err
	}
	companyID, err := company.ParseCompanyID(input.CompanyID)
	if err != nil {
		return nil, err
	}
	user, err := user.NewUser(input.Username, input.Email, hashedPassword, companyID)
	if err != nil {
		return nil, err
	}
	if exists, _ := u.userService.Exists(user); exists {
		return nil, &shared.ConflictError{Resource: input.Email}
	}
	if err := u.userRepository.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Login(input *LoginInput) (auth.Token, error) {
	user, err := u.userRepository.FindByEmail(user.Email(input.Email))
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", &shared.NotFoundError{Resource: input.Email}
	}
	rawPassword, err := auth.NewRawPassword(input.RawPassword)
	if err != nil {
		return "", &shared.UnauthorizedError{}
	}
	if !u.passwordService.Match(rawPassword, user.HashedPassword()) {
		return "", &shared.UnauthorizedError{}
	}
	token, err := u.tokenService.Generate()
	if err != nil {
		return "", err
	}
	return token, nil
}
