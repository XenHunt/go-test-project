package refresh_route

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/uptrace/bun"
)

type ReqRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func New(log *slog.Logger, db *bun.DB, ctx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ReqRes
		if err := render.DecodeJSON(r.Body, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info("Getted access and refresh tokens for new pair")
	}
}
