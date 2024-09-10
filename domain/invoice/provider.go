package invoice

import (
	"sync"
	"upsider-coding-test/shared"

	"github.com/google/wire"
)

var (
	ivcFct     *invoiceFactory
	ivcFctOnce sync.Once

	InvoiceProviderSet wire.ProviderSet = wire.NewSet(
		ProvideInvoiceFactory,
		wire.Bind(new(InvoiceFactory), new(*invoiceFactory)),
	)
)

func ProvideInvoiceFactory(clock shared.Clock) *invoiceFactory {
	ivcFctOnce.Do(func() {
		ivcFct = &invoiceFactory{
			clock: clock,
		}
	})
	return ivcFct
}
