package shared

import (
	"sync"

	"github.com/google/wire"
)

var (
	clck     *clock
	clckOnce sync.Once

	SharedProviderSet wire.ProviderSet = wire.NewSet(
		ProvideClock,
		wire.Bind(new(Clock), new(*clock)),
	)
)

func ProvideClock() *clock {
	clckOnce.Do(func() {
		clck = &clock{}
	})
	return clck
}
