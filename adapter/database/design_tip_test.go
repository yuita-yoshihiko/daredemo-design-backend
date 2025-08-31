package database_test

import (
	"testing"
	"time"

	"github.com/yuita-yoshihiko/daredemo-design-backend/adapter/database"
	"github.com/yuita-yoshihiko/daredemo-design-backend/infrastructure/db"
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
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
		want *models.DesignTip
	}{
		{
			name: "idで抽出した単一のDesignTipのデータが取得できる",
			args: args{
				id: 1,
			},
			want: &models.DesignTip{
				ID:        1,
				Title:     "テストタイトル",
				Guidance:  "テストガイダンス",
				URL:       "https://test.com",
				Media:     "book",
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
			testutils.AssertResponse(t, got, tt.want, models.DesignTip{})
		})
	}
}
