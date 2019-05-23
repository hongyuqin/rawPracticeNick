package main

import (
	"fmt"
	"time"
)

func switchTest(x interface{}) {
	//var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
func selectTest() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // sleep one second
		timeout <- true
	}()
	//ch := make (chan int)
	select {
	case <-timeout:
		fmt.Println("timeout!")
	}
}
func selectTestCh() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}
}
func main() {
	//switchTest(true)
	//selectTest()
	selectTestCh()
}
