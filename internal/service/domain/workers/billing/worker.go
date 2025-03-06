package billing

import (
	"context"
	"log"

	"github.com/recovery-flow/subscriptions-tracker/internal/service"
)

type Worker struct {
	id      int
	tasks   <-chan BillingTask
	service *service.Service
}

func NewWorker(id int, tasks <-chan BillingTask, svc *service.Service) *Worker {
	return &Worker{
		id:      id,
		tasks:   tasks,
		service: svc,
	}
}

func (w *Worker) Start(ctx context.Context) {
	log.Printf("Worker %d started", w.id)
	for {
		select {
		case task, ok := <-w.tasks:
			if !ok {
				log.Printf("Worker %d: task channel closed", w.id)
				return
			}
			// Обработка задачи списания
			w.processTask(ctx, task)
		case <-ctx.Done():
			log.Printf("Worker %d: context cancelled", w.id)
			return
		}
	}
}

func (w *Worker) processTask(ctx context.Context, task BillingTask) {
	err := w.service.Domain.MadeTransaction(ctx, task.UserID)
	if err != nil {
		log.Printf("Worker %d: error processing transaction for user %s: %v", w.id, task.UserID, err)
	} else {
		log.Printf("Worker %d: transaction processed successfully for user %s", w.id, task.UserID)
	}
}
