package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/snowman-mh/go-sample/src/handler/presenter"
	"github.com/snowman-mh/go-sample/src/registry"
	"github.com/snowman-mh/go-sample/src/usecase/input"
	"github.com/snowman-mh/go-sample/src/usecase/usecase"
)

type User struct {
	repository registry.Repository
	presenter  presenter.User
}

func NewUser(repo registry.Repository) User {
	return User{
		repository: repo,
		presenter:  presenter.NewUser(),
	}
}

func (userHandler User) createUserDomain() usecase.UserDomain {
	return usecase.UserDomain{
		User: userHandler.repository.NewUser(),
	}
}

func (userHandler User) Add(w http.ResponseWriter, r *http.Request) {
	var in input.UserAdd
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		userHandler.presenter.Encode(w, http.StatusBadRequest, nil, err)
		return
	}

	domain := userHandler.createUserDomain()
	userUsecase := usecase.NewUser(domain)
	out, err := userUsecase.Add(in)
	if err != nil {
		userHandler.presenter.Encode(w, http.StatusInternalServerError, nil, err)
		return
	}

	userHandler.presenter.Encode(w, http.StatusOK, out, nil)
}

func (userHandler User) Fetch(w http.ResponseWriter, r *http.Request) {
	var in input.UserFetch
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		userHandler.presenter.Encode(w, http.StatusBadRequest, nil, err)
		return
	}
	in.ID = id

	domain := userHandler.createUserDomain()
	userUsecase := usecase.NewUser(domain)
	out, err := userUsecase.Fetch(in)
	if err != nil {
		userHandler.presenter.Encode(w, http.StatusInternalServerError, nil, err)
		return
	}

	userHandler.presenter.Encode(w, http.StatusOK, out, nil)
}
