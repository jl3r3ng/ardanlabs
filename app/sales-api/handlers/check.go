package handlers

import (
	"context"
	"github.com/pkg/errors"
	"gitlab.com/tleuzhan13/service/foundation/web"
	"log"
	"math/rand"
	"net/http"
)

type check struct {
	log *log.Logger
}

func (c check) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if n := rand.Intn(100); n%2 == 0 {
		return web.NewRequestError(errors.New("trusted error"), http.StatusNotFound)
	}
	status := struct {
		Status string
	}{
		Status: "service ready",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
