package handlers

import (
	"os"

	"go.uber.org/zap"

	"github.com/dev-addict/go-service/app/services/sales-api/handlers/v1/healthcheckgrp"
	"github.com/dev-addict/go-service/business/web/v1/mid"
	"github.com/dev-addict/go-service/zarf/web"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *web.App {
	r := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log))

	r.Handle("GET", "/healthcheck", healthcheckgrp.HealthCheck)

	return r
}
