package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/snowman-mh/go-sample/src/handler/handler"
	"github.com/snowman-mh/go-sample/src/registry"
)

func main() {
	r := chi.NewRouter()
	repo := registry.NewRepository()

	user := handler.NewUser(repo)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", user.Add)
		r.Get("/", user.Fetch)
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
