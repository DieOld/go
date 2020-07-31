package eulersProject

import (
	"./solutions"
	"testing"
)

func Test_1(t *testing.T) {
	/*
		Function should return sum of lcm of 5 and 3
	*/
	sum := solutions.Euler1(10)
	if sum != 23 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 23)
	}
}