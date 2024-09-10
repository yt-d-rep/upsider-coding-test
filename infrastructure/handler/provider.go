package handler

import (
	"sync"
	"upsider-base/usecase"

	"github.com/google/wire"
)

var (
	userHdl     *userHandler
	userHdlOnce sync.Once

	ivcHdl     *invoiceHandler
	ivcHdlOnce sync.Once

	HandlerProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserHandler,
		ProvideInvoiceHandler,
		wire.Bind(new(UserHandler), new(*userHandler)),
		wire.Bind(new(InvoiceHandler), new(*invoiceHandler)),
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

func ProvideInvoiceHandler(iUsc usecase.InvoiceUsecase) *invoiceHandler {
	ivcHdlOnce.Do(func() {
		ivcHdl = &invoiceHandler{
			invoiceUsecase: iUsc,
		}
	})
	return ivcHdl
}
