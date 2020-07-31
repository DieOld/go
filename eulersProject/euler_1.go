/*
	Find sum of numbers that lcm of 3 and 5
*/

package main

import (
	"fmt"
)

func euler1(limit int) int {
	sum := 0
	for i := 3; i < limit; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(euler1(1000))
}
