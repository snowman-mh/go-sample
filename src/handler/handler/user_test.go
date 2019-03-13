package handler

import (
	"net/http"
	"testing"

	"github.com/snowman-mh/go-sample/src/handler/presenter"
	"github.com/snowman-mh/go-sample/src/registry"
	"github.com/snowman-mh/go-sample/src/usecase/input"
	"github.com/stretchr/testify/assert"
)

var testUser User

func beforeUserTest() {
	testUser = User{
		repository: registry.NewRepositoryMock(),
		presenter:  presenter.NewUser(),
	}
}

func TestUser_Add(t *testing.T) {
	beforeUserTest()

	body := input.UserAdd{
		GivenName:  "given_name",
		FamilyName: "family_name",
	}

	recorder := record(t, "POST", "/users", nil, body, testUser.Add)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestUser_Fetch(t *testing.T) {
	beforeUserTest()

	params := map[string]string{
		"id": "1",
	}

	recorder := record(t, "GET", "/users", params, nil, testUser.Fetch)
	assert.Equal(t, http.StatusOK, recorder.Code)
}
