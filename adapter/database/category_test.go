package database_test

import (
	"testing"
	"time"

	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/database"
	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
	"github.com/yuita-yoshihiko/daredemo-design-backend/testutils"
)

func Test_Category_FetchAll(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/categories/fetch_all")
	dbUtils := db.NewDBUtil(data)
	r := database.NewCategoryRepository(dbUtils)

	tests := []struct {
		name string
		want []*models.Category
	}{
		{
			name: "Categoryのデータが全て取得できる",
			want: []*models.Category{
				{
					ID:        1,
					Name:      "テストカテゴリ1",
					CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        2,
					Name:      "テストカテゴリ2",
					CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:        3,
					Name:      "テストカテゴリ3",
					CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FetchAll(t.Context())
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want)
		})
	}
}
