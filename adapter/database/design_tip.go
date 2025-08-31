package database

import (
	"context"
	"fmt"

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

func (r *designTipRepositoryImpl) Create(ctx context.Context, m *models.DesignTip) (int64, error) {
	const query = `
		INSERT INTO design_tips (
			title, guidance, url, media, created_at, updated_at
		) VALUES (
			:title, :guidance, :url, :media, NOW(), NOW()
		)
		RETURNING id
	`
	var id int64
	rows, err := r.db.ConnectionFromCtx(ctx).NamedQueryContext(ctx, query, m)
	if err != nil {
		return 0, r.db.Error(err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, r.db.Error(err)
		}
	} else {
		return 0, fmt.Errorf("failed to insert design_tip and retrieve id")
	}
	return id, nil
}
