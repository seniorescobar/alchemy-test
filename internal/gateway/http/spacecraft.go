package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
)

type (
	service interface {
		List(context.Context) ([]spacecraft.Spacecraft, error)
		Get(context.Context, uuid.UUID) (spacecraft.Spacecraft, error)
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

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(spacecrafts); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(spacecraft); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func renderError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	fmt.Fprintln(w, msg)
}
