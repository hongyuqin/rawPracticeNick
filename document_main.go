package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

//官方文档
//1.自定义error
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

//2.goroutines
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func testGoroutine() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	fmt.Scanln()
	fmt.Println("done")
}

//3.通道同步
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

//4.测试ticker
func testTicker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

//5.测试时间
func testTime() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))
}

//6.日期格式化
func testTimeFormat() {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)
	//layout 布局，就是模板
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

//7.测试随机数
func testRand() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float64())
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
	//只有这个是会随机变得 其他的运行多少次都一样
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}

//8.测试数字格式化
func testNumFormat() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

//9.命令行参数
func commandline() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

//10.看下空指针会不会panic
/*func testPanic(p *Person) {
	fmt.Println("person is :", p.name)
}*/

//11.测下返回
func testReturn() (er error) {
	return
}

//12.测试下select超时
func testTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

//13.测试下 通过管道传输后，值是否变空
func testChannelVal() {
	str := "hahaha"
	c1 := make(chan string, 1)
	c1 <- str
	fmt.Println(<-c1, "  ", str)
}

//14.反射
func testReflect2() {
	number := 190
	fmt.Println(reflect.TypeOf(number), "  ", reflect.ValueOf(number))
}

//15.空接口断言，来实现反射类型强转
type Employee struct {
	Name string
	Age  int
}

func reflectPrint(v interface{}) {
	empVal, ok := v.(*Employee)
	if ok {
		log.Println(empVal)
	}
}
func testAssert() {
	emp := &Employee{"naonao", 99}
	reflectPrint(emp)
}
func main() {
	testAssert()
	//13
	//testReflect2()
	//testChannelVal()
	//12.
	//testTimeout()
	//1.测试goroutine
	//testGoroutine()
	//2.通道同步
	/*done := make(chan bool, 1)
	go worker(done)
	<-done*/
	//4.测试ticker
	//testTicker()
	//5.测试时间
	//testTime()
	//6.日期格式化
	//testTimeFormat()
	//7.测试随机数
	//testRand()
	//8.测试数字格式化
	//testNumFormat()
	//9.命令行参数
	//commandline()
	//testPanic(nil)
	//fmt.Println("ddss")
}
