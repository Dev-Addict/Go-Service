package healthcheckgrp

import (
	"context"
	"net/http"

	"github.com/dev-addict/go-service/zarf/web"
)

func HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	response := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, response, http.StatusOK)
}
