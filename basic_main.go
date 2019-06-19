package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"reflect"
	"runtime"
	"strings"
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

//?? 指针为什么也不变xxxx 有的
func (p *Person) doSth() {
	str := "hongmin"
	p.name = str
}

//反射
func TestReflect() {
	var a *Person
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name())
	fmt.Println("========")
	fmt.Println(typeOfA.Kind())
}

//性能分析 ??
/*func joinSlice() []string {
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
}*/

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
func testDuration() {
	now := time.Now()
	time.Sleep(1e9)
	fmt.Println(time.Since(now))
}

//uuid生成器
func testReplace() {
	ret := strings.Replace("abc-sss", "-", "", -1)
	ret1 := strings.Replace("abc-sss", "-", "", -1)
	fmt.Println(ret, "   ", ret1)
}

//锁变量 同步协程
var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock() // 上锁
	counter++
	fmt.Println("counter =", counter)
	lock.Unlock() // 解锁
}
func testLock() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock() // 上锁
		c := counter
		lock.Unlock() // 解锁

		runtime.Gosched() // 出让时间片

		if c >= 10 {
			break
		}
	}
}

type Person struct {
	name string
}

func testValueDeliver(p *Person) {
	p.name = "yibei"
}
func testSuffix() {
	str := "abc,"
	str = str[:len(str)-1]
	fmt.Println(str)
}

//测试map 假如不存在，ok就是false吗
func testMap() {
	scene := make(map[string]int)
	scene["route"] = 66
	_, ok := scene["ss"]
	fmt.Println("testMap :", ok)
}

//看下 数组 是不是 slice
func testSlice1() {
	arr := [...]int{1, 2}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(arr), cap(arr), arr)
	switchTest(arr)
}
func testRemoveMap() {
	scene := make(map[string]int)
	scene["route"] = 66

	fmt.Println("before map :", scene)
	for k, _ := range scene {
		if k == "route" {
			delete(scene, "route")
		}
	}
	fmt.Println("after map :", scene)

}

//测试下map ok有没用
func testNilMap() {
	coupleMap := make(map[string](string))
	coupleMap["hongyuqin"] = "hongyibei"
	couple, ok := coupleMap["hongyuqin"]
	if ok {
		fmt.Println(couple)
	}
	//看下会不会报错
	fmt.Println("==》", coupleMap["hongyuqn"])
}

//测下方法接收者
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func testPointerReceiver() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}

//测试reader
func testReader() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

//测试加密
func encryptHmac256(tokenStr string) string {
	h := sha256.New()
	h.Write([]byte(tokenStr))
	signatureByte := h.Sum(nil)
	signatureStr := fmt.Sprintf("%x", signatureByte)
	signatureStrUpper := strings.ToUpper(signatureStr)
	return signatureStrUpper
}

//测试ToString
type retryParam struct {
	SourcePlatForm string
	RetryStatus    string
	RetryNum       int32
	Limit          int32
}

func (this *retryParam) String() string {
	return "SourcePlatForm:" + this.SourcePlatForm + ",RetryStatus:" + this.RetryStatus + ",RetryNum:" + string(this.RetryNum) + ",Limit:" + string(this.Limit)
}

//测试多个defer会不会执行
func testMultiDefer() {
	fmt.Println("testMultiDefer")
	HandlePanic(nil, func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("painc111")
			}
		}()
		fmt.Println("hello world")

		panic("panic 了")
	})

}
func HandlePanic(ctx context.Context, f func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic222")
		}
	}()
	f()
}
func main() {
	//测试多个defer
	testMultiDefer()
	time.Sleep(time.Second)
	/*rp := &retryParam{SourcePlatForm:"dsfdfs"}
	fmt.Println(rp)*/
	//fmt.Println(encryptHmac256("20181115Royce"))
	//测试下指针作为接收参数
	/*p := &Person{"hongyuqin"}
	p.doSth()
	fmt.Println(p)*/
	//测试反射
	//TestReflect()
	//testReader()
	//testPointerReceiver()
	//testNilMap()
	//testReplace()
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
	//testSlice()
	//testDuration()
	//testLock()
	/*p := &Person{
		name:"hongyuqin",
	}
	testValueDeliver(p)
	fmt.Println(p)*/
	//testSuffix()
	//testMap()//只是当做set用。。。
	//testSlice1()
	//testRemoveMap()

}
