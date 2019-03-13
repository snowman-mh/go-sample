package usecase

import (
	"testing"

	"github.com/snowman-mh/go-sample/src/domain/mock"
	"github.com/snowman-mh/go-sample/src/usecase/input"
	"github.com/stretchr/testify/assert"
)

var testUser User

func beforeUserTest() {
	domain := UserDomain{
		User: mock.NewUser(),
	}
	testUser = User{domain: domain}
}

func TestUser_Add(t *testing.T) {
	beforeUserTest()

	in := input.UserAdd{
		GivenName:  "given_name",
		FamilyName: "family_name",
	}

	out, err := testUser.Add(in)
	if err != nil {
		t.Error(err)
	}
	assert.NotEqual(t, nil, out.UserModel)
}

func TestUser_Fetch(t *testing.T) {
	beforeUserTest()

	in := input.UserFetch{
		ID: uint64(1),
	}

	out, err := testUser.Fetch(in)
	if err != nil {
		t.Error(err)
	}
	assert.NotEqual(t, nil, out.UserModel)
}
