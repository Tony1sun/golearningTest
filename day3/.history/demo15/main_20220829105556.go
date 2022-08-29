package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串的替换
	str1 := strings.Replace("golangandjavagogo", "go", "golang", -1)
	fmt.Println(str1)

	// 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)

	// 将字符串的字母进行大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))
}
