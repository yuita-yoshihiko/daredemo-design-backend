package usecase

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/converter"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/repository"
)

type DesignTipUseCase interface {
	FetchWithCategories(context.Context, int64) (*converter.DesignTipOutput, error)
}

type designTipUseCaseImpl struct {
	repo      repository.DesignTipRepository
	converter converter.DesignTipConverter
}

func NewDesignTipUseCase(
	r repository.DesignTipRepository,
	c converter.DesignTipConverter,
) DesignTipUseCase {
	return &designTipUseCaseImpl{
		repo:      r,
		converter: c,
	}
}

func (u *designTipUseCaseImpl) FetchWithCategories(ctx context.Context, id int64) (*converter.DesignTipOutput, error) {
	m, err := u.repo.FetchWithCategories(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.ToDesignTipOutput(m), nil
}
