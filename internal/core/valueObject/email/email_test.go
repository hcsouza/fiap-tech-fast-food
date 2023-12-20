package email_test

import (
	"testing"

	. "github.com/hcsouza/fiap-tech-fast-food/internal/core/valueObject/email"
	"github.com/stretchr/testify/assert"
)

func TestEmail(t *testing.T) {
	t.Run("should return true when email is valid", func(t *testing.T) {
		email := Email("test@gmail.com")

		assert.True(t, email.IsValid())
	})

	t.Run("should return false when email is invalid", func(t *testing.T) {
		email := Email("testgmail.com")

		assert.False(t, email.IsValid())
	})
}
