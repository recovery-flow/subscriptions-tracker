package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/handlers"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/tokens/identity"
)

func Run(ctx context.Context, svc *service.Service) {
	r := chi.NewRouter()

	r.Use(
		httpkit.CtxMiddleWare(
			handlers.CtxLog(svc.Log),
			handlers.CtxDomain(svc.Domain),
			handlers.CtxConfig(svc.Config),
		),
	)

	authMW := tokens.AuthMdl(svc.Config.JWT.AccessToken.SecretKey)
	roleGrant := tokens.IdentityMdl(svc.Config.JWT.AccessToken.SecretKey, identity.Admin, identity.SuperUser)

	r.Route("/re-news/sub-tracker", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Route("/types", func(r chi.Router) {
					r.Post("/", nil)
					r.Put("/{id}", nil)
				})
				r.Route("/plans", func(r chi.Router) {
					r.Post("/", nil)
					r.Put("/{id}", nil)
				})

				r.Route("/subscriptions", func(r chi.Router) {
					r.Use(authMW)
					r.Route("/{id}", func(r chi.Router) {
						r.Get("/", nil)
						r.Put("/", nil)
					})
				})
			})

			r.Route("/private", func(r chi.Router) {
				r.Use(roleGrant)
				r.Route("/types", func(r chi.Router) {
					r.Post("/", handlers.SubscriptionTypeCreate)
					r.Put("/{id}", handlers.SubscriptionTypeCreate)
				})
				r.Route("/plans", func(r chi.Router) {
					r.Post("/", handlers.SubscriptionPlanCreate)
					r.Put("/{id}", nil)
				})

				r.Route("/subscriptions", func(r chi.Router) {
					r.Route("/{id}", func(r chi.Router) {
						r.Get("/", nil)
						r.Put("/", nil)
					})
				})

				r.Route("/payment_method", func(r chi.Router) {
					r.Get("/", nil)
					r.Put("/", nil)
				})

				r.Route("/transactions", func(r chi.Router) {
					r.Get("/", nil)
					r.Post("/", nil)
				})

				r.Route("/billing_schedules", func(r chi.Router) {
					r.Get("/", nil)
					r.Post("/", nil)
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, svc.Config.Server.Port, r, svc.Log)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, svc.Log)
}
