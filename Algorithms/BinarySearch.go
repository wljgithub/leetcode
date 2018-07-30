package main

import (
	"fmt"
)

func BinarySearch(arr []int, key int) (index int) {
	l, h := 0, len(arr)-1
	//fmt.Println(l, h,mid)
	for l <= h {
		mid := l + (h-l)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			h = mid - 1

		} else {
			l = mid + 1
		}

	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(BinarySearch(arr, 8))
}
