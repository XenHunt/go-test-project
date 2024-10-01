package login_router

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	db_module "github.com/XenHunt/go-test-project/internal/database"
	manager "github.com/XenHunt/go-test-project/internal/token_manager"
	"github.com/go-chi/render"
	"github.com/uptrace/bun"
)

type Request struct {
	GUID string `json:"guid"`
}

type Response struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
func New(log *slog.Logger, db *bun.DB, ctx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := ReadUserIP(r)
		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Info("Logging in with guid - ", req.GUID, " and ip - ", ip)
		ac, err := manager.CreateAccessToken(req.GUID, ip)
		if err != nil {
			log.Error("Couldn't make access token")
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}

		ref := manager.CreateRefreshToken(req.GUID, ip)

		if err := db_module.AddToken(db, ref, *ctx); err != nil {
			log.Error("Couldn't add token to data base")
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		render.JSON(w, r, Response{
			AccessToken:  ac,
			RefreshToken: ref,
		})

		w.WriteHeader(http.StatusOK)
	}
}
