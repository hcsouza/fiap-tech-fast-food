package category

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {
	t.Run("should return true when category is Lanche", func(t *testing.T) {
		category := Category("Lanche")

		assert.True(t, category.IsValid())
	})
	t.Run("should return true when category is Bebida", func(t *testing.T) {
		category := Category("Bebida")

		assert.True(t, category.IsValid())
	})
	t.Run("should return true when category is Acompanhamento", func(t *testing.T) {
		category := Category("Acompanhamento")

		assert.True(t, category.IsValid())
	})
	t.Run("should return true when category is Sobremesa", func(t *testing.T) {
		category := Category("Sobremesa")

		assert.True(t, category.IsValid())
	})
	t.Run("should return false when category is unkown", func(t *testing.T) {
		category := Category("Não mapeada")

		assert.False(t, category.IsValid())
	})
}
