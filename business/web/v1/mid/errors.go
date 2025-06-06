package mid

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	v1 "github.com/dev-addict/go-service/business/web/v1"
	"github.com/dev-addict/go-service/zarf/web"
)

func Errors(log *zap.SugaredLogger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := handler(ctx, w, r); err != nil {
				log.Errorw("ERROR", "trace_id", web.GetTraceID(ctx), "message", err)

				var er v1.ErrorResponse
				var status int

				switch {
				case v1.IsRequestError(err):
					reqErr := v1.GetRequestError(err)
					er = v1.ErrorResponse{
						Error: reqErr.Error(),
					}
					status = reqErr.Status
				default:
					er = v1.ErrorResponse{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				if err := web.Respond(ctx, w, er, status); err != nil {
					return err
				}

				if web.IsShutdown(err) {
					return err
				}
			}

			return nil
		}

		return h
	}

	return m
}
