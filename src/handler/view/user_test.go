package view

import (
	"testing"

	"github.com/snowman-mh/go-sample/src/domain/model"
	"github.com/snowman-mh/go-sample/src/usecase/output"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	id := uint64(1)
	givenName := "given_name"
	familyName := "family_name"
	out := &output.UserAdd{
		UserModel: &model.User{
			ID:         id,
			GivenName:  givenName,
			FamilyName: familyName,
		},
	}

	user := NewUser(out)
	assert := assert.New(t)
	assert.Equal(id, user.ID)
	assert.Equal(givenName, user.GivenName)
	assert.Equal(familyName, user.FamilyName)
}
