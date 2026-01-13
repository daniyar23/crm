package usecase

import (
	"context"

	"github.com/daniyar23/crm/internal/feature/feature1/events"
)

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
					_ = companyService.DeleteCompaniesByUser(ctx, uint(ev.UserID))

				}
			}
		}
	}()
}
