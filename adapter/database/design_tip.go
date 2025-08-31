package database

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/repository"
)

type designTipRepositoryImpl struct {
	db db.DBUtils
}

func NewDesignTipRepository(db db.DBUtils) repository.DesignTipRepository {
	return &designTipRepositoryImpl{db: db}
}

func (r *designTipRepositoryImpl) Fetch(ctx context.Context, id int64) (*models.DesignTip, error) {
	const query = "SELECT * FROM design_tips WHERE id = $1"
	var designTip models.DesignTip
	if err := r.db.ConnectionFromCtx(ctx).GetContext(ctx, &designTip, query, id); err != nil {
		return nil, r.db.Error(err)
	}
	return &designTip, nil
}

func (r *designTipRepositoryImpl) Create(ctx context.Context, m *models.DesignTip) error {
	const query = "INSERT INTO design_tips (name, email) VALUES ($1, $2)"
	_, err := r.db.ConnectionFromCtx(ctx).NamedExecContext(ctx, query, m)
	return r.db.Error(err)
}

func (r *designTipRepositoryImpl) Update(ctx context.Context, m *models.DesignTip) error {
	const query = "UPDATE design_tips SET name = $1, email = $2 WHERE id = $3"
	_, err := r.db.ConnectionFromCtx(ctx).NamedExecContext(ctx, query, m)
	return r.db.Error(err)
}
