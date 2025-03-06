package billing

import (
	"context"
	"log"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/service"
)

const WorkerCount = 10

func Run(ctx context.Context, svc *service.Service, interval time.Duration) {
	tasks := make(chan BillingTask, 100)

	for i := 0; i < WorkerCount; i++ {
		worker := NewWorker(i+1, tasks, svc)
		go worker.Start(ctx)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker pool: context cancelled, shutting down")
			close(tasks)
			return
		case <-ticker.C:
			now := time.Now().UTC()
			schedule, err := svc.Domain.SelectSchedule(ctx, false, now)
			if err != nil {
				log.Printf("Worker pool: error selecting schedule: %v", err)
				continue
			}
			if schedule == nil {
				log.Println("Worker pool: no due billing schedule found")
				continue
			}
			task := BillingTask{
				Schedule: *schedule,
				UserID:   schedule.UserID,
			}
			select {
			case tasks <- task:
				log.Printf("Worker pool: dispatched billing task for user %s", task.UserID)
			case <-ctx.Done():
				close(tasks)
				return
			}
		}
	}
}
