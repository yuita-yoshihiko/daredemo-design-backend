package database

import (
	"context"

	"github.com/lib/pq"
	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
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

func (r *designTipRepositoryImpl) FetchAllWithCategories(ctx context.Context) ([]*custom.DesignTipWithCategories, error) {
	const designTipQuery = `
		SELECT * FROM design_tips
		ORDER BY id
	`
	var designTips []*custom.DesignTipWithCategories
	if err := r.db.ConnectionFromCtx(ctx).SelectContext(ctx, &designTips, designTipQuery); err != nil {
		return nil, r.db.Error(err)
	}
	if len(designTips) == 0 {
		return []*custom.DesignTipWithCategories{}, nil
	}
	ids := make([]int64, len(designTips))
	for i, dt := range designTips {
		ids[i] = dt.ID
	}
	designTipMap := make(map[int64]*custom.DesignTipWithCategories, len(designTips))
	for _, dt := range designTips {
		designTipMap[dt.ID] = dt
	}
	const categoryQuery = `
		SELECT
			dtc.design_tip_id,
			c.*
		FROM categories c
		INNER JOIN design_tip_categories dtc ON c.id = dtc.category_id
		WHERE dtc.design_tip_id = ANY($1)
		ORDER BY dtc.design_tip_id, c.id
	`
	var rows []struct {
		DesignTipID int64 `db:"design_tip_id"`
		models.Category
	}
	if err := r.db.ConnectionFromCtx(ctx).SelectContext(ctx, &rows, categoryQuery, pq.Array(ids)); err != nil {
		return nil, r.db.Error(err)
	}
	for _, rw := range rows {
		if dt := designTipMap[rw.DesignTipID]; dt != nil {
			dt.Categories = append(dt.Categories, rw.Category)
		}
	}
	return designTips, nil
}
