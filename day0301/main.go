package main

import "fmt"

// 运算符
func main() {
	s1 := "中国"
	s2 := "中国"
	s3 := string("中国")
	fmt.Println(s1 == s2)
	fmt.Println(s1 == s3)

	n1 := 3
	n2 := 2
	fmt.Println(n2 > n1)
	fmt.Println(n2 << 10)
	fmt.Println(n2 >> 10)
}
