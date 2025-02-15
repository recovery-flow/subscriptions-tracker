package handlers

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service"
)

type Handlers struct {
	svc *service.Service
}

func NewHandlers(svc *service.Service) *Handlers {
	return &Handlers{svc: svc}
}
