package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "strconv"

	"github.com/go-chi/chi"
	"github.com/veramarycela/api-medica/pkg/paciente"
	"github.com/veramarycela/api-medica/pkg/response"
)

type PacienteRouter struct {
	Repository paciente.Repository
}

func (ur *PacienteRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var p paciente.Paciente
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// p.Password = ""
	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), p.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"paciente": p})
}

func (ur *PacienteRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pacientes, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"paciente": pacientes})
}

func (ur *PacienteRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	// 	return
	// }

	ctx := r.Context()
	p, err := ur.Repository.GetOne(ctx, idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"paciente": p})
}

func (ur *PacienteRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	// 	return
	// }

	var p paciente.Paciente
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Update(ctx, idStr, p)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (ur *PacienteRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	// 	return
	// }

	ctx := r.Context()
	err := ur.Repository.Delete(ctx, idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (ur *PacienteRouter) Routes() http.Handler {
	r := chi.NewRouter()
	// TODO: add routes.
	r.Get("/", ur.GetAllHandler)
	r.Post("/", ur.CreateHandler)
	r.Get("/{id}", ur.GetOneHandler)
	r.Put("/{id}", ur.UpdateHandler)
	r.Delete("/{id}", ur.DeleteHandler)

	return r
}
