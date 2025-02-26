package domain

import (
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra"
	"github.com/sirupsen/logrus"
)

type Domain interface {
}

type domain struct {
	Infra *infra.Infra
	log   *logrus.Logger
}

func NewDomain(infra *infra.Infra, log *logrus.Logger) (Domain, error) {
	return &domain{
		Infra: infra,
		log:   log,
	}, nil
}
