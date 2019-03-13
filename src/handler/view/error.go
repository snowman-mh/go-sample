package view

import (
	"github.com/pkg/errors"
)

type Error struct {
	Message string `json:"message"`
}

var (
	ErrBadRequest     = errors.New("Bad Request")
	ErrInternalServer = errors.New("Internal Server Error")
	ErrUnknown        = errors.New("Unknown Error")
)

func NewBadRequestError(message string) Error {
	return Error{
		Message: errors.Wrap(ErrBadRequest, message).Error(),
	}
}

func NewInternalServerError(message string) Error {
	return Error{
		Message: errors.Wrap(ErrInternalServer, message).Error(),
	}
}

func NewUnknownError(message string) Error {
	return Error{
		Message: errors.Wrap(ErrUnknown, message).Error(),
	}
}
