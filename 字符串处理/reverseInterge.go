package main

import "fmt"

func reverse(x int) int {
	// 把数字转为数组
	// 如果是0，不要放进数组
	var arr []int
	var flag bool
	for x >= 1 || x <= -1 {
		tail := x % 10
		if tail != 0 {
			flag = true
			arr = append(arr, tail)
		} else if flag {
			arr = append(arr, tail)

		}
		x /= 10
	}
	fmt.Println(arr)
	// 从数组中取出元素
	sum := 0
	count := 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i] * power(10, count)
		count += 1
	}
	if sum > 2<<31-1 || sum < -2<<31 {
		return 0
	}
	return sum
}

func power(num, p int) int {
	temp := 1
	for i := 0; i < p; i++ {
		temp *= num
	}
	return temp

}
func main() {
	fmt.Println(reverse(-12))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(24077))
	fmt.Println(reverse(1563847412), 2<<30-1)
	fmt.Println(reverse(2 << 1))
}
