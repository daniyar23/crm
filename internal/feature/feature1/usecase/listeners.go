package usecase

import "context"

func StartUserDeletedListener(
	ctx context.Context,
	events *EventBus,
	companyService CompanyService,
) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case e := <-events.Subscribe():
				switch ev := e.(type) {
				case UserDeletedEvent:
					_ = companyService.DeleteCompaniesByUser(ctx, ev.UserID)
				}
			}
		}
	}()
}
