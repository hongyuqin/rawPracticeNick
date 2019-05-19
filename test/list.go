package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("fist")
	l.PushFront(67)
	//1.遍历链表
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
