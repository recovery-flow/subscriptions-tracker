package cli

import (
	"context"
	"sync"

	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/workers/billing"
)

func runServices(ctx context.Context, srv *service.Service, wg *sync.WaitGroup) {
	var (
		billingSchedules = make(chan struct{})
	)
	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { billing.Run(ctx, srv, billingSchedules) })
	<-billingSchedules

	run(func() { api.Run(ctx, srv) })
}
