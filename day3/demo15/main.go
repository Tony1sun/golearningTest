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

	// 将字符串左右两边的空格去掉
	fmt.Println(strings.TrimSpace("    go and java   "))

	// 将字符串左右两边指定的字符去掉
	fmt.Println(strings.Trim("~golang~", "~"))

	// 将字符串左边指定的字符去掉
	fmt.Println(strings.TrimLeft("~golang~", "~"))

	// 将字符串右边指定的字符去掉
	fmt.Println(strings.TrimRight("~golang~", "~"))

	// 判断字符串是否以指定的字符串开头
	fmt.Println(strings.HasPrefix("http://www.baidu.com", "http"))

	// 判断字符串是否以指定的字符串结尾
	fmt.Println(strings.HasSuffix("http://www.baidu.com", "com"))

}
