package interfaces

import "github.com/NorthDice/ReflectDiary/internal/entity"

type UserRepository interface {
	Save(user *entity.User) (int, error)
	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int) error
}
