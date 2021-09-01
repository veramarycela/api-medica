package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()

	ur := &PacienteRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}
	r.Mount("/pacientes", ur.Routes())
	return r
}
