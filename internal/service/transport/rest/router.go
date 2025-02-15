package rest

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/transport/rest/handlers"
)

func Run(ctx context.Context, svc *service.Service) {
	r := chi.NewRouter()
	h := handlers.NewHandlers(svc)

	r.Route("/re-flow", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Post("/refresh", h.CreatePlan)
			})
		})
	})
}
