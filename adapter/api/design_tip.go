package api

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase"
)

type DesignTipApi struct {
	uc usecase.DesignTipUseCase
}

func NewDesignTipApi(uc usecase.DesignTipUseCase) *DesignTipApi {
	return &DesignTipApi{uc: uc}
}

func (a *DesignTipApi) FetchWithCategories(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		slog.ErrorContext(r.Context(), "Invalid id format", "id", idStr, "error", err.Error())
		WriteJSON(w, http.StatusBadRequest, ErrFailedToFetch)
		return
	}
	u, err := a.uc.FetchWithCategories(r.Context(), id)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed to fetch design_tip", "id", chi.URLParam(r, "id"), "error", err.Error())
		WriteJSON(w, http.StatusInternalServerError, ErrFailedToFetch)
		return
	}
	WriteJSON(w, http.StatusOK, u)
}
