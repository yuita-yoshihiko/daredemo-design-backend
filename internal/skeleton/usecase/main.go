package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yuita-yoshihiko/daredemo-design-backend/internal/skeleton"
)

const Template = `package usecase

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/models/graphql"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/converter"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/repository"
)

type {{ .Upper }}UseCase interface {
	Fetch(context.Context, int64) (*graphql.{{ .Upper }}, error)
	Create(context.Context, graphql.{{ .Upper }}CreateInput) (*graphql.{{ .Upper }}, error)
	Update(context.Context, graphql.{{ .Upper }}UpdateInput) (*graphql.{{ .Upper }}, error)
}

type {{ .Lower }}UseCaseImpl struct {
	repo repository.{{ .Upper }}Repository
	converter  converter.{{ .Upper }}Converter
}

func New{{ .Upper }}UseCase(
	r repository.{{ .Upper }}Repository,
	c converter.{{ .Upper }}Converter,
) {{ .Upper }}UseCase {
	return &{{ .Lower }}UseCaseImpl{
		repo: r,
		converter:  c,
	}
}

func (u *{{ .Lower }}UseCaseImpl) Fetch(ctx context.Context, id int64) (*graphql.{{ .Upper }}, error) {
	m, err := u.repo.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.To{{ .Upper }}Output(m), nil
}

func (u *{{ .Lower }}UseCaseImpl) Create(ctx context.Context, input graphql.{{ .Upper }}CreateInput) (*graphql.{{ .Upper }}, error) {
	m, err := u.converter.To{{ .Upper }}ModelForCreate(input)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Create(ctx, m); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, m.ID)
}

func (u *{{ .Lower }}UseCaseImpl) Update(ctx context.Context, input graphql.{{ .Upper }}UpdateInput) (*graphql.{{ .Upper }}, error) {
	m, err := u.converter.To{{ .Upper }}ModelForUpdate(input)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Update(ctx, m); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, m.ID)
}

`

func main() {
	log.Println("開始")
	list := skeleton.GetNameList()
	for _, m := range list {
		err := skeleton.TemplateExport(m, func(name string) (*os.File, error) {
			return createGoFile(name)
		}, Template)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("完了")
}

func createGoFile(name string) (*os.File, error) {
	filePath := filepath.Join("usecase", fmt.Sprintf("%v.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		log.Printf("ファイルはすでに存在します。 filePath = %s\n", filePath)
		return nil, nil
	}
	log.Printf("usecaseを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
