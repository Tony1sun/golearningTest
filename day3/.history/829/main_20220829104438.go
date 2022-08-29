package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 字符串转整数
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)

	// 整数转字符串
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	// 统计一个字符串有几个指定的子串
	count := strings.Count("golangandjavaga", "ga")
	fmt.Println(count)

	// 不区分大小写的字符串比较
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)

	// 返回字串在字符串第一次出现的索引值，如果没有返回-1
	num2 := strings.Index("golangandjavaga", "ga")
	fmt.Println(num2)

}
