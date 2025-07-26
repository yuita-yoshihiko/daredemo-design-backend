package router

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/api"
)

func NewRouter() *chi.Mux {
	// slogの初期化
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		// ミドルウェア
		r.Use(middleware.Logger)
	})
	setupHealthRoutes(r)

	return r
}

func setupHealthRoutes(r chi.Router) {
	healthApi := api.NewHealthApi()
	r.Get("/health", healthApi.FetchHealth)
}
