package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ppob-api.go/ppob-api/entity"
	"ppob-api.go/ppob-api/repository"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserServiceImpl{UserRepository: userRepository}

func TestCheckPhoneNumber(t *testing.T) {
	user := entity.User{
		Nohandphone: "0812345678910",
	}
	userRepository.Mock.On("FindByPhone", "0812345678910").Return(user)

	isNumberExist := userService.CheckPhoneNumber("0812345678910")
	assert.Equal(t, true, isNumberExist)
}

func TestRegister(t *testing.T) {
	oid := primitive.NewObjectID()

	user := entity.User{
		Nohandphone: "0812345678910",
		Macaddress:  "00:00:00:00",
		Id:          oid,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	userRepository.Mock.On("Create", user).Return(user)

	assert.Equal(t, user, user)
}

func TestGenerateOTP(t *testing.T) {
	oid := primitive.NewObjectID()

	user := entity.User{
		Nohandphone: "0812345678917",
		Id:          oid,
		Otp:         "091231",
		OtpExpire:   time.Now().Add(2 * time.Minute),
	}

	userRepository.Mock.On("Update", user).Return(user)

	assert.Equal(t, user, user)
}
