package rolling

import (
	"fmt"
	"testing"
)

func TestPlus(t *testing.T) {
	counter := Empty

	counter = counter.Plus([]int{10, 1, 2})
	counter = counter.Plus([]int{10, 2, 4})
	counter = counter.Plus([]int{10, 1, 0})

	fmt.Println(counter.ErrPer)
	if counter.ErrPer != 25 {
		t.Error("calc percent error")
	}
}
