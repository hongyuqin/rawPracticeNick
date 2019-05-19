package main

import "fmt"

func main() {
	fmt.Println("begin")
	defer recover()
	panic("crash")
	fmt.Println("recover")
}
