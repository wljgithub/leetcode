package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	move := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums)
		if nums[i] == 0 {
			move++
		} else if move != 0 {
			nums[i-move] = nums[i]
			nums[i] = 0
		}
	}
	fmt.Println(nums)
}
func main() {
	moveZeroes([]int{0, 1, 0, 3, 12, 3123, 442, 0, 112})
}
