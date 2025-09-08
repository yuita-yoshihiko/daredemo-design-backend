package repository

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
)

type CategoryRepository interface {
	FetchAll(context.Context) ([]*models.Category, error)
}
