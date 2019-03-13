package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/snowman-mh/go-sample/src/handler/view"
	"github.com/snowman-mh/go-sample/src/usecase/output"
)

type User struct {
}

func NewUser() User {
	return User{}
}

func (User) Encode(w http.ResponseWriter, code int, out output.User, err error) {
	w.WriteHeader(code)
	switch code {
	case http.StatusOK:
		json.NewEncoder(w).Encode(view.NewUser(out))
	case http.StatusBadRequest:
		json.NewEncoder(w).Encode(view.NewBadRequestError(err.Error()))
	case http.StatusInternalServerError:
		json.NewEncoder(w).Encode(view.NewInternalServerError(err.Error()))
	default:
		json.NewEncoder(w).Encode(view.NewUnknownError(err.Error()))
	}
}
