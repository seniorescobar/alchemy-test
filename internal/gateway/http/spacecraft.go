package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/seniorescobar/alchemy-test/internal/domain/spacecraft"
)

type (
	service interface {
		List(context.Context) ([]spacecraft.Spacecraft, error)
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
