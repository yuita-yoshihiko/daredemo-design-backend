package repository

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
)

type DesignTipRepository interface {
	Fetch(context.Context, int64) (*models.DesignTip, error)
	Create(context.Context, *models.DesignTip) (int64, error)
}
