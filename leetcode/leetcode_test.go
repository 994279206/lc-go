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
