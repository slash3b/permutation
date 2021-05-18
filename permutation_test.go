package permutation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringsPermutation(t *testing.T) {
	result := Strings([]string{"a", "b", "c"})

	assert.Equal(t, len(result), 6)
	assert.Equal(t, result, [][]string{
		{"a", "b", "c"},
		{"a", "c", "b"},
		{"b", "a", "c"},
		{"b", "c", "a"},
		{"c", "a", "b"},
		{"c", "b", "a"},
	})
}

func TestIntsPermutation(t *testing.T) {
	result := Ints([]int{1, 2, 3})

	assert.Equal(t, len(result), 6)
	assert.Equal(t, result, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	})
}
