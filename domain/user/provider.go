package user

import (
	"sync"

	"github.com/google/wire"
)

var (
	userSvc     *userService
	userSvcOnce sync.Once

	UserProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserService,
		wire.Bind(new(UserService), new(*userService)),
	)
)

func ProvideUserService(uRepo UserRepository) *userService {
	userSvcOnce.Do(func() {
		userSvc = &userService{
			uRepo: uRepo,
		}
	})
	return userSvc
}
