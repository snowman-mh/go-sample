package input

import "github.com/snowman-mh/go-sample/src/domain/model"

type UserAdd struct {
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
}

func (in UserAdd) Model() model.User {
	return model.User{
		GivenName:  in.GivenName,
		FamilyName: in.FamilyName,
	}
}

type UserFetch struct {
	ID uint64
}

func (in UserFetch) Model() model.User {
	return model.User{
		ID: in.ID,
	}
}
