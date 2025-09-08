package converter

import "github.com/yuita-yoshihiko/daredemo-design-backend/models"

type CategoryOutput struct {
	Name string `json:"name"`
}

type CategoryConverter interface {
	ToCategoryOutput(*models.Category) *CategoryOutput
	ToCategoryOutputs([]*models.Category) []*CategoryOutput
}

type categoryConverterImpl struct {
}

func NewCategoryConverter() CategoryConverter {
	return &categoryConverterImpl{}
}

func (c *categoryConverterImpl) ToCategoryOutput(input *models.Category) *CategoryOutput {
	return &CategoryOutput{
		Name: input.Name,
	}
}

func (c *categoryConverterImpl) ToCategoryOutputs(inputs []*models.Category) []*CategoryOutput {
	outputs := make([]*CategoryOutput, len(inputs))
	for i, input := range inputs {
		outputs[i] = c.ToCategoryOutput(input)
	}
	return outputs
}
