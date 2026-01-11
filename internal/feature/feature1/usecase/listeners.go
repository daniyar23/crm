package usecase

import (
	"context"

	"your_project/internal/events"
)

type CompanyService interface {
	DeleteCompaniesByUser(ctx context.Context, userID int) error
}

func RunListeners(
	ctx context.Context,
	eventBus events.EventBus,
	companyService CompanyService,
) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case e := <-eventBus.Subscribe():
				switch ev := e.(type) {

				case events.UserDeleted:
					_ = companyService.DeleteCompaniesByUser(ctx, ev.UserID)

				}
			}
		}
	}()
}
