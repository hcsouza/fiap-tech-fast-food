package category

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {
	t.Run("should return true when category is Lanche", func(t *testing.T) {
		category, err := NewCategory("Lanche")

		assert.Equal(t, category, "lanche")
		assert.Nil(t, err)
	})
	t.Run("should return true when category is Bebida", func(t *testing.T) {
		category, err := NewCategory("Bebida")

		assert.Equal(t, category, "bebida")
		assert.Nil(t, err)
	})
	t.Run("should return true when category is Acompanhamento", func(t *testing.T) {
		category, err := NewCategory("Acompanhamento")

		assert.Equal(t, category, "acompanhamento")
		assert.Nil(t, err)
	})
	t.Run("should return true when category is Sobremesa", func(t *testing.T) {
		category, err := NewCategory("Sobremesa")

		assert.Equal(t, category, "sobremesa")
		assert.Nil(t, err)
	})
	t.Run("should return false when category is unkown", func(t *testing.T) {
		category, err := NewCategory("Não mapeada")

		assert.Equal(t, category, "")
		assert.NotNil(t, err)
	})
}
