package main

import "fmt"

func reverse(x int) int {
	// 把数字转为数组
	// 如果是0，不要放进数组
	var arr []int
	for x>=10{
		tail:=x%10
		if tail!=0{
			arr = append(arr,tail)
		}
		x/=10
	}

	// 反转数组
	left,right:=0,len(arr)-1
	for left<right{
		arr[left],arr[right] = arr[right],arr[left]
		left++
		right--
	}

	// 从数组中取出元素
	produce := 1
	count:=0
	for i:=len(arr)-1;i>=0;i--{
		produce *= arr[i] *power(10,count)
		count+=1
	}
	return produce
}

func power(num,p int) int{
	temp :=1
	for i:=0;i<p;i++{
		temp *=num
	}
	return temp

}
func main()  {
	fmt.Println(reverse(123))
}
