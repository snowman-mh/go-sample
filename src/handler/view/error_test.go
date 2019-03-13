package view

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestError(t *testing.T) {
	errType := ErrBadRequest.Error()
	message := "message"

	err := NewBadRequestError(message)
	assert := assert.New(t)
	assert.Contains(err.Message, errType)
	assert.Contains(err.Message, message)
}

func TestNewInternalServerError(t *testing.T) {
	errType := ErrInternalServer.Error()
	message := "message"

	err := NewInternalServerError(message)
	assert := assert.New(t)
	assert.Contains(err.Message, errType)
	assert.Contains(err.Message, message)
}

func TestNewUnknownError(t *testing.T) {
	errType := ErrUnknown.Error()
	message := "message"

	err := NewUnknownError(message)
	assert := assert.New(t)
	assert.Contains(err.Message, errType)
	assert.Contains(err.Message, message)
}
