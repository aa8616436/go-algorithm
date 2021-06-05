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

		assert.Equal(t, 2, count)
	})

	t.Run("findContentChildren", func(t *testing.T) {
		g := []int{10, 9, 8, 7}
		s := []int{5, 6, 7, 8}
		count := findContentChildren(g, s)

		assert.Equal(t, 2, count)
	})

	t.Run("findContentChildren2", func(t *testing.T) {
		g := []int{1, 2}
		s := []int{1, 2, 3}
		count := findContentChildren2(g, s)

		assert.Equal(t, 2, count)
	})
}

func Test135(t *testing.T) {
	t.Run("candy", func(t *testing.T) {
		ints := []int{1, 0, 2} // 1,0,1

		i := candy(ints)

		assert.Equal(t, 5,i)
	})

	t.Run("candy", func(t *testing.T) {
		ints := []int{1,2,87,87,87,2,1}  // 0,1,2,0,2,1,0

		i:= candy(ints)
		assert.Equal(t, 13,i)
	})

}

func Test435(t *testing.T) {
	t.Run("eraseOverlapIntervals", func(t *testing.T) {
		i := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
		r := eraseOverlapIntervals(i)
		assert.Equal(t, 1,r)
	})

	t.Run("eraseOverlapIntervals", func(t *testing.T) {
		i := [][]int{{1, 2}, {1, 2}, {1, 2}}
		r := eraseOverlapIntervals(i)
		assert.Equal(t, 2,r)
	})
}
