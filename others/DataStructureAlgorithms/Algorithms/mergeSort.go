package main

import (
	"fmt"
)

func merge(a, b []int) []int {
	// 因为两个数组都是排序好的，所以用双指针来比较即可（假设两个数组都是相同长度的）
	indexA, indexB := 0, 0
	temp := []int{}
	for {
		// 如果其中一个子数组元素去完，则将另一个数组添加到temp的末尾
		if indexA >= len(a) {
			temp = append(temp, b[indexB:]...)
			break
		}
		if indexB >= len(b) {
			temp = append(temp, a[indexA:]...)
			break
		}

		if a[indexA] < b[indexB] {
			temp = append(temp, a[indexA])
			indexA++
		} else {
			temp = append(temp, b[indexB])
			indexB++
		}

	}
	return temp
}

func main() {
	// 两个子数组长度相同的情况下
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8, 10}
	c := merge(a, b)
	fmt.Println(c)

	// 长度不同的情况下
	a = []int{1, 3, 5, 7, 9, 11}
	b = []int{2, 4, 6, 8, 10}
	c = merge(a, b)
	fmt.Println(c)
}
