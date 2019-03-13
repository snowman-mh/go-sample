package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_FullName(t *testing.T) {
	givenName := "hoge"
	familyName := "fuga"
	userModel := User{
		ID:         1,
		GivenName:  givenName,
		FamilyName: familyName,
	}

	expected := fmt.Sprintf("%s %s", givenName, familyName)
	actual := userModel.FullName()
	assert.Equal(t, expected, actual)
}
