package Greedy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test455(t *testing.T) {
	t.Run("findContentChildren", func(t *testing.T) {
		g := []int{1, 2}
		s := []int{1, 2, 3}
		count := findContentChildren(g, s)

		assert.Equal(t, 2,count)
	})

	t.Run("findContentChildren", func(t *testing.T) {
		g := []int{10,9,8,7}
		s := []int{5,6,7,8}
		count := findContentChildren(g, s)

		assert.Equal(t, 2,count)
	})

	t.Run("findContentChildren2", func(t *testing.T) {
		g := []int{1, 2}
		s := []int{1, 2, 3}
		count := findContentChildren2(g, s)

		assert.Equal(t, 2,count)	})
}
