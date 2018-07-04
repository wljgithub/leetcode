package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	str := strconv.Itoa(num)
	l := len(str)
	if l > 1 {

	}
	arr := make([]int, l)
	fmt.Println(str)
	for i, v := range str {
		arr[i] = int(v) - 48
	}
	fmt.Println(arr)
	return 1
}

func main() {
	// for i, v := range "123" {
	// 	fmt.Println(i, v)
	// }
	fmt.Println(addDigits(123))
	// fmt.Println(string(rune(123)))

}
