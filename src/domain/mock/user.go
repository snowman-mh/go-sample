package mock

import (
	"github.com/snowman-mh/go-sample/src/domain/model"
	"github.com/snowman-mh/go-sample/src/domain/repository"
)

type userMock struct {
}

func NewUser() repository.User {
	return userMock{}
}

func (userMock) Add(*model.User) error {
	return nil
}

func (userMock) Fetch(*model.User) error {
	return nil
}
