package repository //Repository — «как хранить и доставать»
// За что отвечает
// 		ЧТО можно сделать с хранилищем
// 		НЕ КАК именно

// Handler не должен знать,
// Service не должен знать,
// ГДЕ и КАК хранятся пользователи.
import (
	"context"

	"github.com/daniyar23/crm/internal/domain"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
}
