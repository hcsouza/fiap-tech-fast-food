package category

import "slices"

type Category string

func (category Category) IsValid() bool {
	categoryList := []string{
		"Acompanhamento",
		"Bebida",
		"Lanche",
		"Sobremesa",
	}

	return slices.Contains(categoryList, string(category))
}
