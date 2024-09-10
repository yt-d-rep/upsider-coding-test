package usecase

import (
	"sync"
	"upsider-base/domain/auth"
	"upsider-base/domain/invoice"
	"upsider-base/domain/user"

	"github.com/google/wire"
)

var (
	userUsc     *userUsecase
	userUscOnce sync.Once

	ivcUsc     *invoiceUsecase
	ivcUscOnce sync.Once

	UsecaseProviderSet wire.ProviderSet = wire.NewSet(
		ProvideUserUsecase,
		ProvideInvoiceUsecase,
		wire.Bind(new(UserUsecase), new(*userUsecase)),
		wire.Bind(new(InvoiceUsecase), new(*invoiceUsecase)),
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

func ProvideInvoiceUsecase(
	ivcFct invoice.InvoiceFactory,
	ivcRepo invoice.InvoiceRepository,
) *invoiceUsecase {
	ivcUscOnce.Do(func() {
		ivcUsc = &invoiceUsecase{
			invoiceFactory:    ivcFct,
			invoiceRepository: ivcRepo,
		}
	})
	return ivcUsc
}
