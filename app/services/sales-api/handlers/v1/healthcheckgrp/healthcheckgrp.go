package healthcheckgrp

import (
	"encoding/json"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(response)
}
