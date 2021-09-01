package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/veramarycela/api-medica/internal/data"
)

func New() http.Handler {
	r := chi.NewRouter()

	pr := &PacienteRouter{
		Repository: &data.PacienteRepository{
			Data: data.New(),
		},
	}

	r.Mount("/pacientes", pr.Routes())

	mr := &MedicoRouter{
		Repository: &data.MedicoRepository{
			Data: data.New(),
		},
	}

	r.Mount("/medicos", mr.Routes())

	er := &EspecialidadRouter{
		Repository: &data.EspecialidadRepository{
			Data: data.New(),
		},
	}

	r.Mount("/especialidades", er.Routes())

	hr := &HistoriaRouter{
		Repository: &data.HistoriaRepository{
			Data: data.New(),
		},
	}

	r.Mount("/historias", hr.Routes())
	return r
}
