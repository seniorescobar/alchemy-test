package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
)

type (
	service interface {
		List(context.Context, ...spacecraft.Filter) ([]spacecraft.Spacecraft, error)
		Get(context.Context, int) (spacecraft.Spacecraft, error)
		Create(context.Context, spacecraft.Spacecraft) error
		Update(context.Context, spacecraft.Spacecraft) error
		Delete(context.Context, int) error
	}

	SpacecraftGateway struct {
		svc service
	}
)

func NewSpacecraftGateway(svc service) *SpacecraftGateway {
	return &SpacecraftGateway{
		svc: svc,
	}
}

func (g *SpacecraftGateway) List(w http.ResponseWriter, r *http.Request) {
	var filters []spacecraft.Filter
	for _, filter := range spacecraft.AllowedFilters {
		if v := r.FormValue(filter); v != "" {
			filters = append(filters, spacecraft.Filter{Key: filter, Value: v})
		}
	}

	spacecrafts, err := g.svc.List(r.Context(), filters...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(spacecrafts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (g *SpacecraftGateway) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		renderError(w, http.StatusBadRequest, "invalid spacecraft ID provided")
		return
	}

	spacecraft, err := g.svc.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(spacecraft); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (g *SpacecraftGateway) Create(w http.ResponseWriter, r *http.Request) {
	var spacecraft spacecraft.Spacecraft
	if err := json.NewDecoder(r.Body).Decode(&spacecraft); err != nil {
		renderError(w, http.StatusBadRequest, "invalid spacecraft provided")
		return
	}
	defer r.Body.Close()

	if err := g.svc.Create(r.Context(), spacecraft); err != nil {
		if isValidationErr(err) {
			renderError(w, http.StatusBadRequest, err.Error())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	renderSuccess(w, http.StatusCreated)
}

func (g *SpacecraftGateway) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		renderError(w, http.StatusBadRequest, "invalid spacecraft ID provided")
		return
	}

	var spacecraft spacecraft.Spacecraft
	if err := json.NewDecoder(r.Body).Decode(&spacecraft); err != nil {
		renderError(w, http.StatusBadRequest, "invalid spacecraft provided")
		return
	}
	defer r.Body.Close()

	spacecraft.ID = id
	if err := g.svc.Update(r.Context(), spacecraft); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	renderSuccess(w, http.StatusOK)
}

func (g *SpacecraftGateway) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		renderError(w, http.StatusBadRequest, "invalid spacecraft ID provided")
		return
	}

	if err := g.svc.Delete(r.Context(), id); err != nil {
		if errors.Is(err, spacecraft.ErrSpacecraftNotFound) {
			renderError(w, http.StatusNotFound, err.Error())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	renderSuccess(w, http.StatusOK)
}

func isValidationErr(err error) bool {
	var valErr *spacecraft.ValidationErr
	return errors.As(err, &valErr)
}
