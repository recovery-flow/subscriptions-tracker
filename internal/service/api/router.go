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
					r.Route("/{type_id}", func(r chi.Router) {
						r.Get("/", handlers.GetSubscriptionType)
					})
					r.Get("/", handlers.GetSubscriptionTypeActive)
				})

				r.Route("/plan", func(r chi.Router) {
					r.Route("/{plan_id}", func(r chi.Router) {
						r.Get("/", handlers.GetSubscriptionPlan)
					})
				})

				r.Route("/payments", func(r chi.Router) {
					r.Use(authMW)
					r.Post("/", handlers.CreatePayment)
					//r.Get("/", handlers.GetPayments)
					r.Route("/{payment_id}", func(r chi.Router) {
						r.Get("/", handlers.GetPaymentMethod)
						r.Delete("/", handlers.DeletePaymentMethod)
					})
				})

				r.Route("/subscriptions", func(r chi.Router) {
					r.Route("/status", func(r chi.Router) {
						r.Use(authMW)
						r.Post("/activate", handlers.UserSubscriptionActivate)
						r.Patch("/deactivate", handlers.UserSubscriptionDeactivate)
					})
					r.Get("/{user_id}", handlers.UserSubscriptionGet)
				})
			})

			r.Route("/private", func(r chi.Router) {
				r.Use(roleGrant)
				r.Route("/types", func(r chi.Router) {
					r.Post("/", handlers.SubscriptionTypeCreate)
					r.Route("/{id}", func(r chi.Router) {
						r.Put("/", handlers.SubscriptionTypeUpdate)
						r.Post("/activate", handlers.SubscriptionTypeActivate)
						r.Post("/deactivate", handlers.SubscriptionTypeDeactivate)
					})
				})
				r.Route("/plans", func(r chi.Router) {
					r.Post("/", handlers.SubscriptionPlanCreate)
					r.Route("/{id}", func(r chi.Router) {
						r.Put("/", handlers.SubscriptionPlanUpdate)
						r.Post("/activate", handlers.SubscriptionPlanActivate)
						r.Post("/deactivate", handlers.SubscriptionPlanDeactivate)
					})
				})
			})
		})
	})

	server := httpkit.StartServer(ctx, svc.Config.Server.Port, r, svc.Log)

	<-ctx.Done()
	httpkit.StopServer(context.Background(), server, svc.Log)
}
