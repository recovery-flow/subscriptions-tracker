package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/api/handlers"
	"github.com/recovery-flow/tokens"
	"github.com/recovery-flow/tokens/identity"
)

func Run(ctx context.Context, svc *service.Service) {
	r := chi.NewRouter()

	h, err := handlers.NewHandler(svc)
	if err != nil {
		svc.Log.Fatalf("failed to create handlers: %v", err)
		<-ctx.Done()
		return
	}

	authMW := tokens.AuthMdl(svc.Config.JWT.AccessToken.SecretKey)
	roleGrant := tokens.IdentityMdl(svc.Config.JWT.AccessToken.SecretKey, identity.Admin, identity.SuperUser)

}
