package database_test

import (
	"testing"
	"time"

	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/database"
	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models/custom"
	"github.com/yuita-yoshihiko/daredemo-design-backend/testutils"
)

func Test_DesignTip_Fetch(t *testing.T) {
	data := testutils.LoadFixture(t, "testfixtures/design_tips/fetch")
	dbUtils := db.NewDBUtil(data)
	r := database.NewDesignTipRepository(dbUtils)

	type args struct {
		id int64
	}

	tests := []struct {
		name string
		args args
		want *custom.DesignTipWithCategories
	}{
		{
			name: "idで抽出した単一のDesignTipのデータが取得できる",
			args: args{
				id: 1,
			},
			want: &custom.DesignTipWithCategories{
				DesignTip: models.DesignTip{
					ID:        1,
					Title:     "テストタイトル",
					Guidance:  "テストガイダンス",
					URL:       "https://test.com",
					Media:     "book",
					CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Categories: []models.Category{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.FetchWithCategories(t.Context(), tt.args.id)
			if err != nil {
				t.Errorf("error = %v", err)
			}
			testutils.AssertResponse(t, got, tt.want)
		})
	}
}
