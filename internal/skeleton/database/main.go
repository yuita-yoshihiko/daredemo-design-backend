package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yuita-yoshihiko/daredemo-design-backend/internal/skeleton"
)

const Template = `package database

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
	"github.com/yuita-yoshihiko/daredemo-design-backend/usecase/repository"
)

type {{ .Lower }}RepositoryImpl struct {
	db db.DBUtils
}

func New{{ .Upper }}Repository(db db.DBUtils) repository.{{ .Upper }}Repository {
	return &{{ .Lower }}RepositoryImpl{db: db}
}

func (r *{{ .Lower }}RepositoryImpl) Fetch(ctx context.Context, id int64) (*models.{{ .Upper }}, error) {
	const query = "SELECT * FROM {{ .Snake }}s WHERE id = $1"
	var {{ .Lower }} models.{{ .Upper }}
	if err := r.db.ConnectionFromCtx(ctx).GetContext(ctx, &{{ .Lower }}, query, id); err != nil {
		return nil, r.db.Error(err)
	}
	return &{{ .Lower }}, nil
}

func (r *{{ .Lower }}RepositoryImpl) Create(ctx context.Context, m *models.{{ .Upper }}) error {
	const query = "INSERT INTO {{ .Snake }}s (name, email) VALUES ($1, $2)"
	_, err := r.db.ConnectionFromCtx(ctx).NamedExecContext(ctx, query, m)
	return r.db.Error(err)
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
	filePath := filepath.Join("adapter", "database", fmt.Sprintf("%v.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		log.Printf("ファイルはすでに存在します。 filePath = %s\n", filePath)
		return nil, nil
	}
	log.Printf("databaseを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
