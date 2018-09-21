package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	l, h := 0, len(numbers)-1
	for l < h {
		sum := numbers[l]+numbers[h]
		if  sum == target {
			return []int{l, h}
		}else if sum 

	}

}

func main() {

}
