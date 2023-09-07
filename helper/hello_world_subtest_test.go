package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtest(t *testing.T) {
	t.Run("Asep", func(t *testing.T) {
		result := Hello("Asep")
		assert.Equal(t, "Hello Asep", result)
	})
	t.Run("Putra", func(t *testing.T) {
		result := Hello("Putra")
		assert.Equal(t, "Hello Putra", result)
	})
}
