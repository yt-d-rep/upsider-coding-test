// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"sync"
	"upsider-coding-test/domain/invoice"
	"upsider-coding-test/domain/user"
	"upsider-coding-test/infrastructure/handler"
	"upsider-coding-test/infrastructure/middleware"
	"upsider-coding-test/infrastructure/persistent"
	"upsider-coding-test/shared"
	"upsider-coding-test/usecase"
)

// Injectors from wire.go:

func Wire() *HandlerCollection {
	tokenService := middleware.ProvideTokenService()
	interceptor := middleware.ProvideInterceptor(tokenService)
	db := persistent.ProvideDB()
	passwordService := middleware.ProvidePasswordService()
	userRepository := persistent.ProvideUserRepository(db, passwordService)
	userService := user.ProvideUserService(userRepository)
	userUsecase := usecase.ProvideUserUsecase(userService, userRepository, passwordService, tokenService)
	userHandler := handler.ProvideUserHandler(userUsecase)
	clock := shared.ProvideClock()
	invoiceFactory := invoice.ProvideInvoiceFactory(clock)
	invoiceRepository := persistent.ProvideInvoiceRepository(db)
	invoiceUsecase := usecase.ProvideInvoiceUsecase(invoiceFactory, invoiceRepository)
	invoiceHandler := handler.ProvideInvoiceHandler(invoiceUsecase)
	handlerCollection := ProvideHandlerCollection(interceptor, userHandler, invoiceHandler)
	return handlerCollection
}

// wire.go:

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
