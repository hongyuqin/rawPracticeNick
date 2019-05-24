package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	// 准备一个内容被打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用sort包进行排序
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}

	//sort包有封装好的 stringSlice
	names1 := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names1)

	// 遍历打印结果
	for _, v := range names1 {
		fmt.Printf("%s\n", v)
	}
}
