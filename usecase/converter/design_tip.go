package converter

import "github.com/yuita-yoshihiko/daredemo-design-backend/models/custom"

type DesignTipOutput struct {
	Title      string           `json:"title"`
	Guidance   string           `json:"guidance"`
	URL        string           `json:"url"`
	Media      string           `json:"media"`
	Categories []CategoryOutput `json:"categories"`
}

type DesignTipConverter interface {
	ToDesignTipOutput(*custom.DesignTipWithCategories) *DesignTipOutput
}

type designTipConverterImpl struct {
}

func NewDesignTipConverter() DesignTipConverter {
	return &designTipConverterImpl{}
}

func (c *designTipConverterImpl) ToDesignTipOutput(input *custom.DesignTipWithCategories) *DesignTipOutput {
	return &DesignTipOutput{
		Title:    input.Title,
		Guidance: input.Guidance,
		URL:      input.URL,
		Media:    input.Media,
		Categories: func() []CategoryOutput {
			categories := make([]CategoryOutput, len(input.Categories))
			for i, category := range input.Categories {
				categories[i] = CategoryOutput{
					Name: category.Name,
				}
			}
			return categories
		}(),
	}
}
