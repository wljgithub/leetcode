package main

import "fmt"

func removeDuplicates(nums []int) int {
	duplicates := 0
	l, h := 0, 1
	for h < len(nums)-1 {
		if nums[l] == nums[h] {
			duplicates++

		}
		l++
		h++
	}
	return len(nums) - duplicates
}
func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
