package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	strs := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			strs[i-1] = "FizzBuzz"
		case i%3 == 0:
			strs[i-1] = "Fizz"
		case i%5 == 0:
			strs[i-1] = "Buzz"
		default:
			strs[i-1] = strconv.Itoa(i)
		}

	}
	return strs
}

func main() {
	arr := fizzBuzz(15)
	fmt.Println(arr)
}
