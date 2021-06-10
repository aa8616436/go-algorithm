package Pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test167(t *testing.T) {
	t.Run("twoSun", func(t *testing.T) {
		ints := []int{2, 7, 11, 15}
		sum := twoSum(ints, 9)
		assert.Equal(t, []int{1,2},sum)
	})
	t.Run("twoSun", func(t *testing.T) {
		ints := []int{2,3,4}
		sum := twoSum(ints, 6)
		assert.Equal(t, []int{1,3},sum)
	})

	t.Run("twoSun", func(t *testing.T) {
		ints := []int{-1,0}
		sum := twoSum(ints, -1)
		assert.Equal(t, []int{1,2},sum)
	})

}

func Test88(t *testing.T) {
	t.Run("merge", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2,5,6}
		merge(nums1,3,nums2,3)
	})

	t.Run("merge", func(t *testing.T) {
		nums1 := []int{1}
		nums2 := []int{}
		merge(nums1,1,nums2,0)
	})
}
