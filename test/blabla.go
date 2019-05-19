package main

import "fmt"
import "flag"

//输入参数
var skillParam = flag.String("skill", "", "skill to perform")

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	//匿名函数
	//1.100表示对匿名函数的调用
	/*func(data int) {
		fmt.Println("hello", data)
	}(100)
	//2.作为回调函数
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})*/

	//?? 3.使用匿名函数实现操作封装
	flag.Parse()
	fmt.Println(*skillParam)
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

}
