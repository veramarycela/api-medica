package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/go-chi/chi"
	"github.com/veramarycela/api-medica/pkg/historia"
	"github.com/veramarycela/api-medica/pkg/response"
)

type HistoriaRouter struct {
	Repository historia.Repository
}

func (ur *HistoriaRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var e historia.Historia
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
	response.JSON(w, r, http.StatusCreated, response.Map{"historia": e})
}

func (ur *HistoriaRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	historias, err := ur.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"historia": historias})
}

func (ur *HistoriaRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	e, err := ur.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"historia": e})
}

func (ur *HistoriaRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var e historia.Historia
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ur.Repository.Update(ctx, uint(id), e)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

func (ur *HistoriaRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
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

func (ur *HistoriaRouter) Routes() http.Handler {
	r := chi.NewRouter()
	// TODO: add routes.
	r.Get("/", ur.GetAllHandler)
	r.Post("/", ur.CreateHandler)
	r.Get("/{id}", ur.GetOneHandler)
	r.Put("/{id}", ur.UpdateHandler)
	r.Delete("/{id}", ur.DeleteHandler)

	return r
}
