//go:build wireinject

package di

import (
	"sync"
	"upsider-base/domain/user"
	"upsider-base/infrastructure/handler"
	"upsider-base/infrastructure/middleware"
	"upsider-base/infrastructure/persistent"
	"upsider-base/usecase"

	"github.com/google/wire"
)

type HandlerCollection struct {
	Interceptor *middleware.Interceptor
	UserHandler handler.UserHandler
}

var (
	hdl     *HandlerCollection
	hdlOnce sync.Once

	HandlerCollectionProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandlerCollection,
	)
)

func ProvideHandlerCollection(
	intrceptr *middleware.Interceptor,
	uHdl handler.UserHandler,
) *HandlerCollection {
	hdlOnce.Do(func() {
		hdl = &HandlerCollection{
			Interceptor: intrceptr,
			UserHandler: uHdl,
		}
	})
	return hdl
}

func Wire() *HandlerCollection {
	panic(wire.Build(
		// domain
		user.UserProviderSet,
		// usecase
		usecase.UsecaseProviderSet,
		// infrastructure
		HandlerCollectionProviderSet,
		handler.HandlerProviderSet,
		middleware.MiddlewareProviderSet,
		persistent.PersistentProviderSet,
	))
}
