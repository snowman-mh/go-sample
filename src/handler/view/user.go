package view

import (
	"time"

	"github.com/snowman-mh/go-sample/src/usecase/output"
)

type User struct {
	ID         uint64     `json:"id"`
	GivenName  string     `json:"given_name"`
	FamilyName string     `json:"family_name"`
	FullName   string     `json:"full_name"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func NewUser(out output.User) User {
	return User{
		ID:         out.User().ID,
		GivenName:  out.User().GivenName,
		FamilyName: out.User().FamilyName,
		FullName:   out.User().FullName(),
		CreatedAt:  out.User().CreatedAt,
		UpdatedAt:  out.User().UpdatedAt,
	}
}
