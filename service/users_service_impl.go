package service

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"ppob-api.go/ppob-api/entity"
	"ppob-api.go/ppob-api/helper"
	"ppob-api.go/ppob-api/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) CheckPhoneNumber(phone string) bool {
	user, _ := service.UserRepository.FindByPhone(phone)

	return (entity.User{}) != user
}

func (service *UserServiceImpl) Create(user *entity.User) error {
	helper.Log(logrus.InfoLevel, "registering user")
	registeredUser, _ := service.UserRepository.FindByPhone(user.Nohandphone)

	if registeredUser != (entity.User{}) {
		helper.Log(logrus.WarnLevel, "user try to register more than once")
		return errors.New("user already registered")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	helper.Log(logrus.InfoLevel, "saving user to db")
	err := service.UserRepository.Save(user)

	if err != nil {
		helper.Log(logrus.WarnLevel, "failed to save user")
		return err
	}

	helper.Log(logrus.InfoLevel, "user registered")
	return nil
}

func (service *UserServiceImpl) GenerateOTP(phone string) error {
	helper.Log(logrus.InfoLevel, "generating OTP")
	user, _ := service.UserRepository.FindByPhone(phone)

	if user == (entity.User{}) {
		helper.Log(logrus.WarnLevel, "user not found")
		return errors.New("user not found")
	}

	otp := helper.RandStringBytes(6)

	user.Otp = otp

	err := service.UserRepository.Update(&user)

	if err != nil {
		return err
	}

	return nil
}
