package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunc001(t *testing.T) {
	n := 100
	ans := func001(n)
	assert.Equal(t, 25, ans, "期望值相同")
}
func TestFunc002(t *testing.T) {
	a := []int{0, 1, 2, 2, 3, 3, 4}
	index := func002(a)
	assert.Equal(t, 5, index, "期望值相同")

}
