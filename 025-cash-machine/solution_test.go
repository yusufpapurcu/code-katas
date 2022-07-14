package cashmachine_test

import (
	"fmt"
	"testing"

	cashmachine "github.com/yusufpapurcu/code-katas/025-cash-machine"
)

func TestBreakIntoChange(t *testing.T) {

	t.Run("", func(t *testing.T) {
		fmt.Println(cashmachine.BreakIntoChange(3.45))
	})
}
