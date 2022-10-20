package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
)

type (
	service interface {
		List(context.Context) ([]spacecraft.Spacecraft, error)
		Get(context.Context, uuid.UUID) (spacecraft.Spacecraft, error)
		Create(context.Context, spacecraft.Spacecraft) error
		Update(context.Context, spacecraft.Spacecraft) error
		Delete(context.Context, uuid.UUID) error
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
	spacecrafts, err := g.svc.List(r.Context())
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
	id, err := uuid.Parse(mux.Vars(r)["id"])
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	renderSuccess(w, http.StatusCreated)
}

func (g *SpacecraftGateway) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
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
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		renderError(w, http.StatusBadRequest, "invalid spacecraft ID provided")
		return
	}

	if err := g.svc.Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	renderSuccess(w, http.StatusNoContent)
}
