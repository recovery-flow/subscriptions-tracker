package cli

import (
	"context"
	"sync"

	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api"
	"github.com/recovery-flow/subscriptions-tracker/internal/workers/billing"
	"github.com/recovery-flow/subscriptions-tracker/internal/workers/listener"
)

func runServices(ctx context.Context, svc *service.Service, wg *sync.WaitGroup) {
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

	run(func() { billing.Run(ctx, svc, billingSchedules) })
	<-billingSchedules

	run(func() { listener.Listen(ctx, svc) })

	run(func() { api.Run(ctx, svc) })

}
