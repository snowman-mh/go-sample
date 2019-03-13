package registry

import (
	"github.com/snowman-mh/go-sample/src/domain/repository"
	"github.com/snowman-mh/go-sample/src/infra/dao"
)

type Repository interface {
	NewUser() repository.User
}

type repositoryImpl struct {
}

func NewRepository() Repository {
	return repositoryImpl{}
}

func (repositoryImpl) NewUser() repository.User {
	return dao.NewUser()
}
