// 冒泡排序 每次比较两个元素，把最大的元素移动到数组末尾
// 比较n-1趟，最终形成有序数组
package main

import (
	"fmt"
	"time"
)

func bubbleSort(arr []int) []int {
	start := time.Now().Unix()
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println("bubble sort use time:", time.Now().Unix()-start)
	return arr
}

func bubbleSort2(arr []int) []int {
	start := time.Now().Unix()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println("bubble sort 2 use time:", time.Now().Unix()-start)
	return arr
}

func main() {
	arr := []int{}
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}
	bubbleSort(arr)
	for i := 100000; i > 0; i-- {
		arr = append(arr, i)
	}
	fmt.Println(bubbleSort2(arr))
}
