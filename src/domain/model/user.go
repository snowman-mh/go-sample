package model

import (
	"fmt"
	"time"
)

type User struct {
	ID         uint64
	GivenName  string
	FamilyName string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (userModel User) FullName() string {
	return fmt.Sprintf("%s %s", userModel.GivenName, userModel.FamilyName)
}
