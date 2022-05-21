package game_of_three_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	game_of_three "github.com/yusufpapurcu/code-katas/001-game-of-three"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		Given    int
		Expected []int
	}{
		{3, []int{3, 1}},
		{100, []int{100, 99, 33, 11, 12, 4, 3, 1}},
		{-100, []int{-100, -99, -33, -11, -12, -4, -3, -1}},
	}

	for _, tcase := range cases {
		got := game_of_three.Solution(tcase.Given)
		assert.Equal(t, tcase.Expected, got)
	}
}
