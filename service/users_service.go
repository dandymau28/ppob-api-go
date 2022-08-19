package service

import "ppob-api.go/ppob-api/entity"

type UserService interface {
	CheckPhoneNumber(phone string) bool
	Create(user *entity.User) error
	GenerateOTP(phone string) error
}
