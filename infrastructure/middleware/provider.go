package middleware

import (
	"sync"
	"upsider-base/domain/auth"

	"github.com/google/wire"
)

var (
	pswdSvc     *passwordService
	pswdSvcOnce sync.Once

	tknSvc     *tokenService
	tknSvcOnce sync.Once

	intrcptr     *Interceptor
	intrcptrOnce sync.Once

	MiddlewareProviderSet wire.ProviderSet = wire.NewSet(
		ProvidePasswordService,
		ProvideTokenService,
		ProvideInterceptor,
		wire.Bind(new(auth.PasswordService), new(*passwordService)),
		wire.Bind(new(auth.TokenService), new(*tokenService)),
	)
)

func ProvidePasswordService() *passwordService {
	pswdSvcOnce.Do(func() {
		pswdSvc = &passwordService{}
	})
	return pswdSvc
}

func ProvideTokenService() *tokenService {
	tknSvcOnce.Do(func() {
		tknSvc = &tokenService{}
	})
	return tknSvc
}

func ProvideInterceptor(tokenSvc auth.TokenService) *Interceptor {
	intrcptrOnce.Do(func() {
		intrcptr = &Interceptor{
			tokenSvc: tokenSvc,
		}
	})
	return intrcptr
}
