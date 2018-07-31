package main

import (
	"fmt"
)

func stringToInt(s string) []int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	arr := []rune(s)

	strings := make([]int, len(arr))
	for i, v := range arr {
		strings[i] = m[string(v)]
		fmt.Println(string(v))
	}
	// fmt.Println(strings)
	return strings
}

func romanToInt(s string) int {
	nums := stringToInt(s)
	fmt.Println(nums)
	sum := 0
	for i, v := range nums {
		if i == len(nums) {
			sum += v

		} else if v > nums[i+1] {
			sum += v
		} else {
			sum -= v
		}

	}
	return sum

}
func main() {
	// StringToInt("III")
	fmt.Println(romanToInt("III"))
}
