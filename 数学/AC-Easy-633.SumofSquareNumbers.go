package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	l := 0
	h := int(math.Sqrt(float64(c)))
	// fmt.Println(h)
	for l <= h {
		Pow := l*l + h*h
		if Pow == c {
			return true
		} else if Pow > c {
			h -= 1
		} else {
			l += 1
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(2))
}
