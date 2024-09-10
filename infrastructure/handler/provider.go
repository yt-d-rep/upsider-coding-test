package handler

import (
	"sync"
	"upsider-base/usecase"

	"github.com/google/wire"
)

var (
	userHdl     *userHandler
	userHdlOnce sync.Once

	HandlerProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserHandler,
		wire.Bind(new(UserHandler), new(*userHandler)),
	)
)

func ProvideUserHandler(uUsc usecase.UserUsecase) *userHandler {
	userHdlOnce.Do(func() {
		userHdl = &userHandler{
			userUsecase: uUsc,
		}
	})
	return userHdl
}
