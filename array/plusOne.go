package main

import "fmt"

func plusOne(arr []int) []int64 {
	// 把数组拼成数字
	var  sum int64 = 0
	count := 0
	for i := len(arr) - 1; i >= 0; i-- {
		sum += int64(arr[i]) * int64(power(10, count))
		count += 1
	}
	fmt.Println("sum:", sum)
	//+1
	sum += 1


	//根据数的位数，来创建对应大小的数组 -- 如123 数组容量： [x,x,x]
	lenght:=1
	temp:=sum
	for temp>=10 {
		lenght+=1
		temp/=10
	}
	l := make([]int64,lenght)

	//将数字还原为数组
	for i:=lenght-1;i>=0 ;i--  {
		l[i] = sum%10
		sum /= 10
	}

	return l
}
func plusOne2(arr []int) []int {
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

	//数组反转
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
	i := []int{3,2,1}
	i = []int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6}
	fmt.Println(plusOne(i))
	fmt.Println(plusOne2(i))


}
