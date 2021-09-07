package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "strconv"

	"github.com/go-chi/chi"
	"github.com/veramarycela/api-medica/pkg/especialidad"
	"github.com/veramarycela/api-medica/pkg/response"
)

type EspecialidadRouter struct {
	Repository especialidad.Repository
}

func (ur *EspecialidadRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var e especialidad.Especialidad
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Create(ctx, &e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// p.Password = ""
	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), e.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"especialidad": e})
}

func (ur *EspecialidadRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	especialidads, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"especialidad": especialidads})
}

func (ur *EspecialidadRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	// 	return
	// }

	ctx := r.Context()
	e, err := ur.Repository.GetOne(ctx, idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"especialidad": e})
}

func (ur *EspecialidadRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	// 	return
	// }

	var e especialidad.Especialidad
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Update(ctx, idStr, e)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (ur *EspecialidadRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
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

func (ur *EspecialidadRouter) Routes() http.Handler {
	r := chi.NewRouter()
	// TODO: add routes.
	r.Get("/", ur.GetAllHandler)
	r.Post("/", ur.CreateHandler)
	r.Get("/{id}", ur.GetOneHandler)
	r.Put("/{id}", ur.UpdateHandler)
	r.Delete("/{id}", ur.DeleteHandler)

	return r
}
