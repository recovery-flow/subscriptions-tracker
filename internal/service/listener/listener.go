package listener

import (
	"context"
	"encoding/json"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Listener struct {
	brokers []string
}

func NewListener(cfg *config.Config) *Listener {
	return &Listener{
		brokers: cfg.Kafka.Brokers,
	}
}

func (l *Listener) Listen(ctx context.Context, log *logrus.Logger) error {
	for _, tc := range TopicsConfig {
		tc := tc
		go func() {
			r := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        l.brokers,
				Topic:          tc.Topic,
				MinBytes:       1,
				MaxBytes:       10e6,
				CommitInterval: time.Second,
			})
			defer r.Close()

			for {
				m, err := r.ReadMessage(ctx)
				if err != nil {
					if ctx.Err() != nil {
						return
					}
					log.WithField("kafka", err).Errorf("Error reading message from topic %s", tc.Topic)
					continue
				}

				var ie events.InternalEvent
				if err := json.Unmarshal(m.Value, &ie); err != nil {
					log.WithField("kafka", err).Error("Error unmarshalling InternalEvent")
					continue
				}

				if err := tc.Callback(ctx, m, ie); err != nil {
					log.WithField("kafka", err).Errorf("Error processing message from topic %s", tc.Topic)
					if tc.OnError != nil {
						tc.OnError(ctx, m, ie, err)
					}
					continue
				}

				if tc.OnSuccess != nil {
					if err := tc.OnSuccess(ctx, m, ie); err != nil {
						log.WithField("kafka", err).Error("Error in OnSuccess callback")
					}
				}
			}
		}()
	}

	<-ctx.Done()
	log.Info("Producer listener stopped")
	return nil
}
