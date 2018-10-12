package main

//链表实现
import (
	"fmt"
)

//定义错误常量
const (
	ERROR = -1000000001
)

//定义元素类型
type Element int64

//定义节点
type LinkNode struct {
	Data Element   //数据域
	Nest *LinkNode //指针域，指向下一个节点
}

//函数接口
type LinkNoder interface {
	Add(head *LinkNode, Data Element)               //后面添加
	Delete(head *LinkNode, index int)               //删除指定index位置元素
	Insert(head *LinkNode, index int, data Element) //在指定index位置插入元素
	GetLength(head *LinkNode) int                   //获取长度
	Search(head *LinkNode, data Element)            //查询元素的位置
	GetData(head *LinkNode, index int) Element      //获取指定index位置的元素
}

//添加 头结点，数据
func Add(head *LinkNode, data Element) {
	point := head //临时指针
	for point.Nest != nil {
		point = point.Nest //移位
	}
	var node LinkNode  //新节点
	point.Nest = &node //赋值
	node.Data = data

	// if GetLength(head) > 1 {
	// 	Traverse(head)
	// }

}

//删除 头结点 index 位置
func Delete(head *LinkNode, index int) Element {
	//判断index合法性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return ERROR
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest //移位
		}

		if point.Nest.Nest != nil {

			point.Nest = point.Nest.Nest //赋值
			data := point.Nest.Data
			return data
		}
	}
	fmt.Println("delete element failed!")
	return ERROR
}

//插入 头结点 index位置 data元素
func Insert(head *LinkNode, index int, data Element) {
	//检验index合法性
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest //移位
		}
		var node LinkNode //新节点，赋值
		node.Data = data
		node.Nest = point.Nest
		point.Nest = &node
	}
}

//获取长度 头结点
func GetLength(head *LinkNode) int {
	if head == nil {
		return 0
	}

	var length = 1
	point := head
	for point.Nest != nil {
		length++
		point = point.Nest
	}
	return length
}

//搜索 头结点 data元素
func Search(head *LinkNode, data Element) {
	point := head
	index := 0
	for point.Nest != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Nest
			if index > GetLength(head)-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}

func Search2(head *LinkNode, data Element) (ele Element) {

	if head == nil {
		return
	}

	point := head
	for index := 0; index < GetLength(head); index++ {
		if point.Data == data {
			fmt.Printf("the element %d exists,and index :%d", data, index)
			return data
		}
		if point.Nest != nil {

			point = point.Nest
		}
	}
	fmt.Println("element not exists!")
	return
}

//获取data 头结点 index位置
func GetData(head *LinkNode, index int) Element {
	point := head
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return ERROR
	} else {
		for i := 0; i < index; i++ {
			point = point.Nest
		}
		return point.Data
	}
}

//遍历 头结点
func Traverse(head *LinkNode) {
	point := head
	for point.Nest != nil {
		fmt.Println(point.Data)
		point = point.Nest
	}
}

//主函数测试
func main() {

	head := LinkNode{Data: 0, Nest: nil}

	for i := 1; i < 10; i++ {

		Add(&head, Element(i))
		fmt.Println("add element:", i)
	}
	// Search2(&head, Element(9))

	Delete(&head, 8)
	Search2(&head, Element(8))

	Insert(&head, 8, Element(8))
	fmt.Println("\n", GetLength(&head))

}
