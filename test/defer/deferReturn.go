package main

import "fmt"

func testDefer() {
	return
	defer fmt.Println("after return")
}

//测试defer放在return后面是否会执行
func main() {
	testDefer()
}
