package main

import (
	"fmt"
	"runtime"
	"sync"
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

//就是监听 IO 操作，当 IO 操作发生时，触发相应的动作
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
func testSyncMap() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

//查询cpu数量
func findCpuNum() {
	cpuNum := runtime.NumCPU()
	fmt.Print(cpuNum)
}
func testTimer() {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	println("Timer expired")
}
func testSlice() {
	a := make([]int, 2)
	a = append(a, 10)
	a = append(a, 101)
	a = append(a, 100)
	fmt.Println(a)
}
func main() {
	//switchTest(true)
	//selectTest()
	//selectTestCh()
	//testSyncMap()
	//findCpuNum()
	//testTimer()
	testSlice()
}
