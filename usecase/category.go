package usecase

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/converter"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/repository"
)

type CategoryUseCase interface {
	FetchAll(context.Context) ([]*converter.CategoryOutput, error)
}

type categoryUseCaseImpl struct {
	repo      repository.CategoryRepository
	converter converter.CategoryConverter
}

func NewCategoryUseCase(
	r repository.CategoryRepository,
	c converter.CategoryConverter,
) CategoryUseCase {
	return &categoryUseCaseImpl{
		repo:      r,
		converter: c,
	}
}

func (u *categoryUseCaseImpl) FetchAll(ctx context.Context) ([]*converter.CategoryOutput, error) {
	ms, err := u.repo.FetchAll(ctx)
	if err != nil {
		return nil, err
	}
	return u.converter.ToCategoryOutputs(ms), nil
}
