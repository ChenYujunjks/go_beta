package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	// 向链表中添加元素： PushBack 在链表尾部插入元素
	// PushFront 在链表头部插入元素
	e1 := l.PushBack("Go")
	l.PushFront("Hello")
	l.PushBack("World")

	// InsertBefore 在指定元素之前插入新元素
	l.InsertBefore("InsertBefore", e1)

	// 遍历链表，使用 Front() 获取链表的头部，Next() 依次访问后续节点
	fmt.Println("遍历链表：")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l.Remove(e1)

	fmt.Println("删除 e1 后遍历链表：")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
