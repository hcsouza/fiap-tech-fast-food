package category

import (
	"slices"
	"strings"

	coreErrors "github.com/hcsouza/fiap-tech-fast-food/internal/core/errors"
)

type category string

func NewCategory(value string) (string, error) {
	valueAsCategory := category(strings.ToLower(value))

	if !valueAsCategory.IsValid() {
		return "", coreErrors.ErrInvalidCategory
	}

	return valueAsCategory.String(), nil
}

func (category category) String() string {
	return string(category)
}

func (category category) IsValid() bool {
	categoryList := []string{
		"acompanhamento",
		"bebida",
		"lanche",
		"sobremesa",
	}

	return slices.Contains(categoryList, strings.ToLower(string(category)))
}
