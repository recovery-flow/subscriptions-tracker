package listener

import (
	"context"
	"fmt"

	"github.com/recovery-flow/subscriptions-tracker/internal/service"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events"
	"github.com/segmentio/kafka-go"
)

func Listener(ctx context.Context, svc *service.Service) {
	// Создаём Kafka-брокер
	broker := events.NewBroker(svc.Config)

	// Определяем набор топиков, которые хотим читать
	topics := []events.TopicConfig{
		{
			Topic: "users_ban",
			Callback: func(ctx context.Context, m kafka.Message) error {
				// допустим, просто печатаем
				fmt.Printf("Banned user message: key=%s, val=%s\n", m.Key, m.Value)
				return nil
			},
		},
		// Можно добавить другие топики / коллбеки
	}

	// Запускаем чтение
	if err := broker.RunConsumers(ctx, topics); err != nil {
		svc.Log.Errorf("Error running consumers: %v", err)
	}

	// Ожидаем завершения контекста (например, ctrl+c)
	<-ctx.Done()
	svc.Log.Info("Kafka listener stopped")
}
