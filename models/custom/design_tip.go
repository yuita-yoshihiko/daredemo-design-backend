package custom

import (
	"github.com/yuita-yoshihiko/daredemo-design-backend/models"
)

type DesignTipWithCategories struct {
	models.DesignTip
	Categories []models.Category
}
