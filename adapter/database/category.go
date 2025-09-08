package database

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/repository"
)

type categoryRepositoryImpl struct {
	db db.DBUtils
}

func NewCategoryRepository(db db.DBUtils) repository.CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

func (r *categoryRepositoryImpl) FetchAll(ctx context.Context) ([]*models.Category, error) {
	const query = "SELECT * FROM categories"
	var categories []*models.Category
	if err := r.db.ConnectionFromCtx(ctx).SelectContext(ctx, &categories, query); err != nil {
		return nil, r.db.Error(err)
	}
	return categories, nil
}
