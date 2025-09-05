package database

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models/custom"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/repository"
)

type designTipRepositoryImpl struct {
	db db.DBUtils
}

func NewDesignTipRepository(db db.DBUtils) repository.DesignTipRepository {
	return &designTipRepositoryImpl{db: db}
}

func (r *designTipRepositoryImpl) FetchWithCategories(ctx context.Context, id int64) (*custom.DesignTipWithCategories, error) {
	const designTipQuery = `
		SELECT * FROM design_tips WHERE id = $1
	`
	var designTip custom.DesignTipWithCategories
	if err := r.db.ConnectionFromCtx(ctx).GetContext(ctx, &designTip, designTipQuery, id); err != nil {
		return nil, r.db.Error(err)
	}
	const categoryQuery = `
		SELECT c.* FROM categories c
		INNER JOIN design_tip_categories dtc ON c.id = dtc.category_id
		WHERE dtc.design_tip_id = $1
		ORDER BY c.id
	`
	if err := r.db.ConnectionFromCtx(ctx).SelectContext(ctx, &designTip.Categories, categoryQuery, id); err != nil {
		return nil, r.db.Error(err)
	}
	return &designTip, nil
}
