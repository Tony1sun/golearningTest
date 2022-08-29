package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串的替换
	str1 := strings.Replace("golangandjavagogo", "go", "golang", -1)
	fmt.Println(str1)
}
