package repository //Repository — «как хранить и доставать»
// За что отвечает
// 		ЧТО можно сделать с хранилищем
// 		НЕ КАК именно

// Handler не должен знать,
// Service не должен знать,
// ГДЕ и КАК хранятся пользователи.
import (
	"github.com/daniyar23/crm/internal/domain"
)

type UserRepository interface { // Мы объявляем контракт, а не реализацию
	GetUserByID(id uint) (domain.User, error)
	// Очевидно и просто — одна ответственность.
	GetAllUsers() ([]domain.User, error)
	// Для начала CRM этого достаточно.
	CreateUser(user domain.User) (domain.User, error)
	// принимает чистую domain-модель
	// возвращает:
	// 		пользователя (уже с ID)
	// 		ошибку (если БД упала, конфликт и т.п.)
	// Почему возвращаем User, а не ID?
	// Потому что:
	// 		сервису часто нужен объект целиком
	// 		это гибче
	DeleteUser(id uint) error
}
