package handlers

import (
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/dev-addict/go-service/app/services/sales-api/handlers/v1/healthcheckgrp"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", healthcheckgrp.HealthCheck).Methods("GET")

	return r
}
