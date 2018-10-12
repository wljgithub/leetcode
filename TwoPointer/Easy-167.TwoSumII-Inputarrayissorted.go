package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, h := 0, len(numbers)-1
	for l < h {
		sum := numbers[l]+numbers[h]
		if  sum == target {
			return []int{l, h}
		}

	}
	return []int{}
}

func main() {
	a,b:=1,2
	a=a^b
	b=a^b
	a=a^b
	fmt.Println(a,b)
}
