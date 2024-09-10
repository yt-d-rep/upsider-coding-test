package usecase

import (
	"sync"
	"upsider-base/domain/auth"
	"upsider-base/domain/user"

	"github.com/google/wire"
)

var (
	userUsc     *userUsecase
	userUscOnce sync.Once

	UsecaseProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserUsecase,
		wire.Bind(new(UserUsecase), new(*userUsecase)),
	)
)

func ProvideUserUsecase(
	uSvc user.UserService,
	uRepo user.UserRepository,
	pSvc auth.PasswordService,
	tSvc auth.TokenService,
) *userUsecase {
	userUscOnce.Do(func() {
		userUsc = &userUsecase{
			userRepository:  uRepo,
			userService:     uSvc,
			passwordService: pSvc,
			tokenService:    tSvc,
		}
	})
	return userUsc
}
