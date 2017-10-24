package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	t.Run("Database", func(t *testing.T) {
		assert.Equal(t, 3, 1+4)
	})
}
