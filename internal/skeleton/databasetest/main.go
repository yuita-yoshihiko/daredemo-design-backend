package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yuita-yoshihiko/daredemo-design-backend/internal/skeleton"
)

const Template = `package database_test

import (
	"testing"
	"time"

	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/database"
	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
	"github.com/yuita-yoshihiko/daredemo-design-backend/testutils"
)

func Test_{{ .Upper }}_Fetch(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/{{ .Snake }}s/fetch")
	dbUtils := db.NewDBUtil(data)
	r := database.New{{ .Upper }}Repository(dbUtils)

	type args struct {
		id int64
	}

	tests := []struct {
		name string
		args args
		want *models.{{ .Upper }}
	}{
		{
			name: "idで抽出した単一の{{ .Upper }}のデータが取得できる",
			args: args{
				id: 1,
			},
			want: &models.{{ .Upper }}{
				// TODO: 必要なフィールドを追加してください
				ID:        1,
				CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Fetch(t.Context(), tt.args.id)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want, models.{{ .Upper }}{})
		})
	}
}

func Test_{{ .Upper }}_Create(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/{{ .Snake }}s/create")
	dbUtils := db.NewDBUtil(data)
	r := database.New{{ .Upper }}Repository(dbUtils)

	tests := []struct {
		name  string
		input *models.{{ .Upper }}
		want  *models.{{ .Upper }}
	}{
		{
			name: "正常に{{ .Upper }}のデータが作成できる",
			input: &models.{{ .Upper }}{
				// TODO: 必要なフィールドを追加してください
			},
			want: &models.{{ .Upper }}{
				// TODO: 必要なフィールドを追加してください
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := r.Create(t.Context(), tt.input)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			got, err := r.Fetch(t.Context(), id)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			opt1 := testutils.DefaultIgnoreFieldsOptions(models.{{ .Upper }}{})
			opt2 := testutils.GenerateIgnoreUnexportedTypesOptions(models.{{ .Upper }}{})
			testutils.AssertResponseWithOptions(t, got, tt.want, opt1, opt2)
		})
	}
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
	filePath := filepath.Join("adapter", "database", fmt.Sprintf("%v_test.go", name))
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		log.Printf("ファイルはすでに存在します。 filePath = %s\n", filePath)
		return nil, nil
	}
	log.Printf("database_testを作成します。 filePath = %s\n", filePath)
	return os.Create(filePath)
}
