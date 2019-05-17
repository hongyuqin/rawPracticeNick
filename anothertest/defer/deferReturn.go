package main

import "fmt"

func testReturn() {
	return
	defer fmt.Println("hello world")
}
func main() {
	testReturn()
}
