package usecase

import (
	"github.com/snowman-mh/go-sample/src/domain/repository"
	"github.com/snowman-mh/go-sample/src/usecase/input"
	"github.com/snowman-mh/go-sample/src/usecase/output"
)

type User struct {
	domain UserDomain
}

type UserDomain struct {
	User repository.User
}

func NewUser(domain UserDomain) User {
	return User{
		domain: domain,
	}
}

func (user User) Add(in input.UserAdd) (*output.UserAdd, error) {
	userModel := in.Model()
	if err := user.domain.User.Add(&userModel); err != nil {
		return nil, err
	}
	return &output.UserAdd{UserModel: &userModel}, nil
}

func (user User) Fetch(in input.UserFetch) (*output.UserFetch, error) {
	userModel := in.Model()
	if err := user.domain.User.Fetch(&userModel); err != nil {
		return nil, err
	}
	return &output.UserFetch{UserModel: &userModel}, nil
}
