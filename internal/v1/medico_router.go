package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/go-chi/chi"
	"github.com/veramarycela/api-medica/pkg/medico"
	"github.com/veramarycela/api-medica/pkg/response"
)

type MedicoRouter struct {
	Repository medico.Repository
}

func (ur *MedicoRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var m medico.Medico
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &m)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// p.Password = ""
	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), m.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"medico": m})
}

func (ur *MedicoRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	medicos, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"medico": medicos})
}

func (ur *MedicoRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	m, err := ur.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"medico": m})
}

func (ur *MedicoRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var m medico.Medico
	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Update(ctx, uint(id), m)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (ur *MedicoRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = ur.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

func (ur *MedicoRouter) Routes() http.Handler {
	r := chi.NewRouter()
	// TODO: add routes.
	r.Get("/", ur.GetAllHandler)
	r.Post("/", ur.CreateHandler)
	r.Get("/{id}", ur.GetOneHandler)
	r.Put("/{id}", ur.UpdateHandler)
	r.Delete("/{id}", ur.DeleteHandler)

	return r
}
