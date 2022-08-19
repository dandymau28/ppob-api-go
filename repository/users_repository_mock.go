package repository

import (
	"errors"

	"github.com/stretchr/testify/mock"
	"ppob-api.go/ppob-api/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) FindByPhone(phone string) (entity.User, error) {
	arguments := repository.Mock.Called(phone)
	if arguments.Get(0) == nil {
		return entity.User{}, nil
	}
	user := arguments.Get(0).(entity.User)
	return user, nil
}

func (repository *UserRepositoryMock) Save(user *entity.User) error {
	arguments := repository.Mock.Called(user)
	if arguments.Get(0) == nil {
		return errors.New("Failed")
	}
	return nil
}

func (repository *UserRepositoryMock) Update(user *entity.User) error {
	arguments := repository.Mock.Called(user)
	if arguments.Get(0) == nil {
		return errors.New("Failed")
	}
	return nil
}
