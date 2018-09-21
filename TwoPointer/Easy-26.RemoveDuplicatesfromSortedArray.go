package main

import "fmt"

func removeDuplicates(nums []int) int {
<<<<<<< HEAD
	//duplicates := 0
	l, h := 1, 2
	for h < len(nums)-1 {
		if nums[l] != nums[h] {
			//duplicates++
			nums[l],nums[h] = nums[h],nums[l]
			l++;h++
		}else{
			h++
		}
	}
	fmt.Println(nums,"len:",l+1)
	return l+1
=======
	l, h := 1, 2
	for h < len(nums)-1 {
		//fmt.Println(nums)
		if nums[l] != nums[h] {
			nums[l],nums[h] = nums[h],nums[l]
			l++
			h++
			fmt.Println(nums)
		}else {
			h++

		}
	}
	return l
>>>>>>> c122d488de4df6aed424db1954d526d995e1f82f
}
func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
