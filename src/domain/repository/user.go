package repository

import "github.com/snowman-mh/go-sample/src/domain/model"

type User interface {
	Add(*model.User) error
	Fetch(*model.User) error
}
