package router

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/api"
	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/database"
	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/converter"
)

func NewRouter() *chi.Mux {
	// slogの初期化
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		// ミドルウェア
		r.Use(middleware.Logger)

		// ルーティング
		dbInstanse, err := db.InitDB()
		if err != nil {
			panic(err)
		}
		dbUtil := db.NewDBUtil(dbInstanse)
		setupDesignTipRoutes(r, dbUtil)
	})
	setupHealthRoutes(r)

	return r
}

func setupHealthRoutes(r chi.Router) {
	healthApi := api.NewHealthApi()
	r.Get("/health", healthApi.FetchHealth)
}

func setupDesignTipRoutes(r chi.Router, dbUtil db.DBUtils) {
	formUseCase := usecase.NewDesignTipUseCase(
		database.NewDesignTipRepository(dbUtil),
		converter.NewDesignTipConverter(),
	)
	handler := api.NewDesignTipApi(formUseCase)

	r.Get("/design_tips/{id}", handler.FetchWithCategories)
}
