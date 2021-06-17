package BinarySearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test69(t *testing.T) {
	t.Run("mySqrt", func(t *testing.T) {
		sqrt := mySqrt(4)
		assert.Equal(t, 2,sqrt)
	})

	t.Run("mySqrt", func(t *testing.T) {
		sqrt := mySqrt(8)
		assert.Equal(t, 2,sqrt)
	})
}
