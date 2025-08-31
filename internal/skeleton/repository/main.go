package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yuita-yoshihiko/daredemo-design-backend/internal/skeleton"
)

const Template = `package repository

import (
	"context"

	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
)

type {{ .Upper }}Repository interface {
	Fetch(context.Context, int64) (*models.{{ .Upper }}, error)
	Create(context.Context, *models.{{ .Upper }}) error
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
	filePath := filepath.Join("usecase", "repository", fmt.Sprintf("%v.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		log.Printf("ファイルはすでに存在します。 filePath = %s\n", filePath)
		return nil, nil
	}
	log.Printf("Repositoryを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
