package entity

import (
	"time"

	"github.com/snowman-mh/go-sample/src/domain/model"
)

type User struct {
	ID         uint64
	GivenName  string
	FamilyName string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func NewUser(userModel *model.User) User {
	return User{
		ID:         userModel.ID,
		GivenName:  userModel.GivenName,
		FamilyName: userModel.FamilyName,
		CreatedAt:  userModel.CreatedAt,
		UpdatedAt:  userModel.UpdatedAt,
	}
}

func (userEntity User) Write(userModel *model.User) {
	userModel.ID = userEntity.ID
	userModel.GivenName = userEntity.GivenName
	userModel.FamilyName = userEntity.FamilyName
	userModel.CreatedAt = userEntity.CreatedAt
	userModel.UpdatedAt = userEntity.UpdatedAt
}
