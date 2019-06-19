package main

import (
	"fmt"
	"strings"
)

//1.去掉字符串前缀
func testPrefix() {
	fmt.Println(strings.TrimPrefix("osd_112222", "osd_"))
	fmt.Println("osd_1112222"[4:])
}

//2.测试map遍历
func testMapFor() {
	testMap := make(map[string]string)
	testMap["hongyuqin"] = "hongyibei"
	testMap["fangmin"] = "wangxuan"
	for v := range testMap {
		fmt.Println(v)
	}
}

//3.测试slice遍历
func testListFor() {
	list := []int{9, 2, 3, 4, 5}
	for v := range list {
		fmt.Println(v)
	}

}

//4.recover捕获panic
func panicF() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")
}
func recoverTest() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	panicF()
}

//5.枚举类型
type State int

const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Rebooting:
		return "Rebooting"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

type T struct {
	Name  string
	Port  int
	State State
}

func main() {
	//5.
	/*t := T{Name: "example", Port: 6666}
	fmt.Printf("t %+v\n", t)
	*/
	recoverTest()
	//1.str->int
	//a,_ := strconv.Atoi("")
	//fmt.Println(a)
	//testPrefix()
	//testMapFor()
	//testListFor()
}
