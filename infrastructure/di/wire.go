//go:build wireinject

package di

import (
	"sync"
	"upsider-base/domain/invoice"
	"upsider-base/domain/user"
	"upsider-base/infrastructure/handler"
	"upsider-base/infrastructure/middleware"
	"upsider-base/infrastructure/persistent"
	"upsider-base/shared"
	"upsider-base/usecase"

	"github.com/google/wire"
)

type HandlerCollection struct {
	Interceptor    *middleware.Interceptor
	UserHandler    handler.UserHandler
	InvoiceHandler handler.InvoiceHandler
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
	iHdl handler.InvoiceHandler,
) *HandlerCollection {
	hdlOnce.Do(func() {
		hdl = &HandlerCollection{
			Interceptor:    intrceptr,
			UserHandler:    uHdl,
			InvoiceHandler: iHdl,
		}
	})
	return hdl
}

func Wire() *HandlerCollection {
	panic(wire.Build(
		// domain
		user.UserProviderSet,
		invoice.InvoiceProviderSet,
		// usecase
		usecase.UsecaseProviderSet,
		// infrastructure
		HandlerCollectionProviderSet,
		handler.HandlerProviderSet,
		middleware.MiddlewareProviderSet,
		persistent.PersistentProviderSet,
		// shared
		shared.SharedProviderSet,
	))
}
