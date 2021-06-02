package main

import (
	"fmt"
	"strings"
)

var s2 = `我的中国心,
我的中国心,中国心`

func main() {
	// 字符串 和 字符

	s := "a中国"
	c1 := '中'
	alpha := "abcdef"
	fmt.Println(c1)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", c1)

	fmt.Println(s2)
	fmt.Println(len(s2)) // 字节长度

	fmt.Println(strings.Split(s2, ","))
	fmt.Println(strings.Contains(s2, ","))
	fmt.Println(strings.Index(s2, ",")) // 字节位置
	fmt.Println(strings.LastIndex(s2, ","))
	fmt.Println(strings.LastIndex(alpha, "c"))

	// rune 切片
	s3 := []rune(s)
	s3[0] = '美'
	fmt.Printf("s3: %T", s3)
	fmt.Println(string(s3))

}
