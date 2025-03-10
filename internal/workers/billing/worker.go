package billing

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/domain/models"
	"github.com/recovery-flow/subscriptions-tracker/internal/workers/cron"
)

func Run(ctx context.Context, svc *service.Service, sig chan struct{}) {
	cron.Init(svc.Log)
	log := svc.Log.WithField("who", "billing")

	_, err := cron.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(0, 15, 0))),
		gocron.NewTask(func() {
			schedules, err := svc.Domain.SelectSchedules(ctx, false, time.Now().UTC(), string(models.ScheduleBillingStatusPlanned))
			if err != nil {
				log.WithError(err).Error("failed to get schedules")
				return
			}

			if len(schedules) == 0 {
				log.Info("No schedules to process")
				return
			}

			const concurrencyLimit = 10
			sem := make(chan struct{}, concurrencyLimit)
			var wg sync.WaitGroup

			for _, schedule := range schedules {
				wg.Add(1)
				sem <- struct{}{}
				go func(sched models.BillingSchedule) {
					defer wg.Done()
					_, err = svc.Domain.MadeTransaction(ctx, sched.UserID)
					if err != nil {
						log.WithError(err).Errorf("failed to make transaction for user %s", sched.UserID)
					} else {
						log.Infof("Transaction processed successfully for user %s", sched.UserID)
					}
					<-sem
				}(schedule)
			}
			wg.Wait()
		}),
		gocron.WithName("billing"),
	)
	if err != nil {
		panic(fmt.Errorf("failed to initialize daily job: %w", err))
	}
	sig <- struct{}{}

	cron.Start(ctx)
}
