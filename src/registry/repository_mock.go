package registry

import (
	"github.com/snowman-mh/go-sample/src/domain/mock"
	"github.com/snowman-mh/go-sample/src/domain/repository"
)

type repositoryMock struct {
}

func NewRepositoryMock() Repository {
	return repositoryMock{}
}

func (repositoryMock) NewUser() repository.User {
	return mock.NewUser()
}
