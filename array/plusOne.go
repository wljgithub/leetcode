package main

import "fmt"

func plusOne(arr []int) []int {
	// 把数组拼成数字
	sum := 0
	count := 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += arr[i] * power(10, count)
		count += 1
	}
	fmt.Println("sum:", sum)
	//+1
	sum += 1
	//将数字还原为数组
	var l []int
	for sum>=1{
		l = append(l,sum%10)
		sum/=10
	}

	// 数组反转
	left,right:=0,len(l)-1
	for left<right{
		l[left],l[right]=l[right],l[left]
		left+=1
		right-=1
	}
	return l
}

var array []int

func trunNumtoArray(i int) {
	if i >= 10 {
		trunNumtoArray(i / 10)
	}
	array = append(array, i%10)
}
func power(num, power int) int {
	produce := 1
	for i := 0; i < power; i++ {
		produce *= num
	}
	return produce
}
func power2(num, power int) int {
	prodce:=1
	for {
		if power>0{
			prodce*=num
			power--
		}else {
			break
		}
	}
	return prodce
}
//func power3(num,power int)int{
//
//}
func main() {
	i := []int{1,0}
	fmt.Println(plusOne(i))


}
