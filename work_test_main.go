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
func main() {
	//1.str->int
	//a,_ := strconv.Atoi("")
	//fmt.Println(a)
	//testPrefix()
	//testMapFor()
	testListFor()
}
