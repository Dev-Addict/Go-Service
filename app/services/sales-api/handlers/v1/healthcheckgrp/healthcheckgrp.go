package healthcheckgrp

import (
	"context"
	"encoding/json"
	"net/http"
)

func HealthCheck(_ context.Context, w http.ResponseWriter, r *http.Request) error {
	response := struct {
		Status string
	}{
		Status: "OK",
	}

	return json.NewEncoder(w).Encode(response)
}
