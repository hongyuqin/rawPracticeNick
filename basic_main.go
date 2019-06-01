package main

import (
	"fmt"
	"github.com/pkg/profile"
	"reflect"
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

type Person struct {
	name string
	age  int
}

//?? 指针为什么也不变
func (p *Person) doSth() {
	str := "hongmin"
	p.name = str
}

//反射
func testReflect() {
	var a *Person
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name())
	fmt.Println("========")
	fmt.Println(typeOfA.Kind())
}

//性能分析 ??
func joinSlice() []string {
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	// 在main()结束时停止性能分析
	defer stopper.Stop()
	// 分析的核心逻辑
	joinSlice()
	// 让程序至少运行1秒
	time.Sleep(time.Second)
	var arr []string
	for i := 0; i < 100000; i++ {
		// 故意造成多次的切片添加(append)操作, 由于每次操作可能会有内存重新分配和移动, 性能较低
		arr = append(arr, "arr")
	}
	return arr
}

//测试通道
func testCh() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			c1 <- i
		}
	}()
	go func() {
		for i := 10; i < 20; i++ {
			c2 <- i
		}
	}()

	for i := 0; i < 10; i++ {
		v1 := <-c1
		v2 := <-c2
		fmt.Printf("c1 = %d;c2 = %d\n", v1, v2)
	}
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
	/*p := new(Person)
	p.name = "hongyuqin"
	p.age = 26
	fmt.Println("person is :",p)*/
	//testReflect()
	// 开始性能分析, 返回一个停止接口
	//测试通道
	//testCh()

	//findCpuNum()
	//testTimer()
	testSlice()
}
