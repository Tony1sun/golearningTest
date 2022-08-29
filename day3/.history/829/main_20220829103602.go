package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转整数
	num1, _ := strconv.Atoi("666")
	fmt.Println(num1)

	// 整数转字符串
	str1 := strconv.Itoa(88)
	fmt.Println(str1)
}
