package output

import "github.com/snowman-mh/go-sample/src/domain/model"

type User interface {
	User() *model.User
}

type UserAdd struct {
	UserModel *model.User
}

func (out UserAdd) User() *model.User {
	return out.UserModel
}

type UserFetch struct {
	UserModel *model.User
}

func (out UserFetch) User() *model.User {
	return out.UserModel
}
