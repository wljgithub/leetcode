

# 链表

链表是一种非连续的数据结构，有一系列节点组成。每个节点由两个部分组成，分别是存储数据的数据域和指向下一个节点的指针域。由于链表不是按顺序存储，链表在插入时可以到达 O(1)级别的时间复杂度，但是查找需要遍历一遍列表，所以时间复杂度为 O(n).

(单向)链表的go语言实现
```go
type Node struct{
	Data Element    //元素域(可以存放各种类型的值)
	Next *Node      // 指针域
}
```

链表常见的操作

- 在链表尾添加元素
- 指定索引删除元素
- 在指定位置插入元素
- 获取链表长度
- 查询某个元素的位置

```go
type LinkNoder interface{
	Add(head *Node,data Element)   				// 增
	Delete(head *Node,index int)				// 删
	Search(head *Node,data Element)             // 查
	Insert(head *Node,index int,data Element)   // 改
	GetLength(head *Node)
}
```

```go

```