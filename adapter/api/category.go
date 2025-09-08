package api

import (
	"log/slog"
	"net/http"

	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase"
)

type CategoryApi struct {
	uc usecase.CategoryUseCase
}

func NewCategoryApi(uc usecase.CategoryUseCase) *CategoryApi {
	return &CategoryApi{uc: uc}
}

func (a *CategoryApi) FetchAll(w http.ResponseWriter, r *http.Request) {
	us, err := a.uc.FetchAll(r.Context())
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to fetch categories", "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToFetch)
		return
	}
	WriteJSON(w, http.StatusOK, us)
}
