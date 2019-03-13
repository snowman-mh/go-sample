package dao

import (
	"testing"

	"github.com/snowman-mh/go-sample/src/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestUserImpl_Add(t *testing.T) {
	expected := model.User{
		GivenName:  "given名前👻",
		FamilyName: "family名字🤡",
	}

	user := NewUser()
	if err := user.Add(&expected); err != nil {
		t.Error(err)
	}
	actual := model.User{}
	if err := user.Fetch(&actual); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)
	assert.Equal(expected.GivenName, actual.GivenName)
	assert.Equal(expected.FamilyName, actual.FamilyName)
	assert.NotEqual(uint64(0), actual.ID)
}

func TestUserImpl_Fetch(t *testing.T) {
	expected := model.User{
		GivenName:  "given名前😧",
		FamilyName: "family名字🥴",
	}
	user := NewUser()
	if err := user.Add(&expected); err != nil {
		t.Error(err)
	}

	actual := model.User{}
	if err := user.Fetch(&actual); err != nil {
		t.Error(err)
	}
	assert := assert.New(t)
	assert.Equal(expected.GivenName, actual.GivenName)
	assert.Equal(expected.FamilyName, actual.FamilyName)
	assert.NotEqual(uint64(0), actual.ID)
}
