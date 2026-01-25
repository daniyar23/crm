package usecase

import (
	"context"
	"encoding/json"
	"github.com/daniyar23/crm/internal/feature/feature1/events"

	"github.com/segmentio/kafka-go"
)

func RunKafkaListeners(
	ctx context.Context,
	brokers string,
	topic string,
	companyService CompanyService,
) {
	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{brokers},
			Topic:   topic,
			GroupID: "crm-company-cleaner",
		})

		for {
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				return
			}

			var event events.UserDeleted
			if err := json.Unmarshal(msg.Value, &event); err == nil {
				_ = companyService.DeleteCompaniesByUser(ctx, uint(event.UserID))
			}
		}
	}()
}
