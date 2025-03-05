package events

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/recovery-flow/subscriptions-tracker/internal/config"
	"github.com/recovery-flow/subscriptions-tracker/internal/service/infra/events/evebody"
	"github.com/segmentio/kafka-go"
)

type Kafka interface {
	SubscriptionCreated(body evebody.CreateSubscription) error
	RunConsumers(ctx context.Context, topics []TopicConfig) error

	sendMessage(msg kafka.Message) error
}

type broker struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
	cfg    *config.Config
}

type TopicConfig struct {
	Topic    string
	Callback func(ctx context.Context, message kafka.Message) error
}

func NewBroker(cfg *config.Config) Kafka {
	//var reqAcks kafka.RequiredAcks
	//switch cfg.Kafka.RequiredAcks {
	//case "all":
	//	reqAcks = kafka.RequireAll
	//case "1":
	//	reqAcks = kafka.RequireOne
	//case "0":
	//	reqAcks = kafka.RequireNone
	//default:
	//	reqAcks = kafka.RequireAll
	//}

	writer := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Kafka.Brokers...),
		Topic:                  "subscription_transactions",
		Balancer:               &kafka.LeastBytes{}, // Оптимальный балансировщик
		RequiredAcks:           kafka.RequireNone,   // Не ждет подтверждения от брокера
		BatchSize:              1,                   // Отправляет сообщения по одному
		BatchTimeout:           0,                   // Не ждет накопления сообщений
		Async:                  true,                // Асинхронный режим (не ждет завершения отправки)
		AllowAutoTopicCreation: true,                // Если топик не существует, создаст его автоматически
	}

	// Reader: здесь используется DialTimeout
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Kafka.Brokers,
		GroupID: "test",
		Topic:   "test",
	})

	return &broker{
		Reader: reader,
		Writer: writer,
		cfg:    cfg,
	}
}

func (b *broker) sendMessage(msg kafka.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := b.Writer.WriteMessages(ctx, msg); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (b *broker) SubscriptionCreated(body evebody.CreateSubscription) error {
	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal ActivateSubscription event: %w", err)
	}

	msg := kafka.Message{
		Key:   []byte(body.UserID),
		Value: data,
	}

	return b.sendMessage(msg)
}

func (b *broker) RunConsumers(ctx context.Context, topics []TopicConfig) error {
	for _, t := range topics {
		// локальная копия в цикле
		tc := t

		go func() {
			// Создаём отдельный Reader для каждого топика (либо можно переиспользовать один).
			r := kafka.NewReader(kafka.ReaderConfig{
				Brokers:  b.cfg.Kafka.Brokers,
				GroupID:  b.cfg.Kafka.GroupID, // либо другой groupID, если нужно
				Topic:    tc.Topic,
				MinBytes: 1,
				MaxBytes: 10e6, // 10MB
			})
			defer r.Close()

			// Читаем в цикле
			for {
				m, err := r.ReadMessage(ctx)
				if err != nil {
					// если ctx отменен, выходим без ошибки
					if ctx.Err() != nil {
						return
					}
					fmt.Printf("Error reading message from topic %s: %v\n", tc.Topic, err)
					return
				}
				// вызываем коллбек
				if cbErr := tc.Callback(ctx, m); cbErr != nil {
					fmt.Printf("Error processing message from topic %s: %v\n", tc.Topic, cbErr)
					// тут можно логировать, решать, делать ли commit offset и т.д.
				}
			}
		}()
	}

	return nil
}
