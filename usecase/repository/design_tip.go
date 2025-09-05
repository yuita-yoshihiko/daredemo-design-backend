package repository

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/models/custom"
)

type DesignTipRepository interface {
	FetchWithCategories(context.Context, int64) (*custom.DesignTipWithCategories, error)
	FetchAllWithCategories(context.Context) ([]*custom.DesignTipWithCategories, error)
}
