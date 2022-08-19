package repository

import (
	"ppob-api.go/ppob-api/entity"
)

type UserRepository interface {
	FindByPhone(phone string) (entity.User, error)
	Save(user *entity.User) error
	Update(user *entity.User) error
}
