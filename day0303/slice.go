package main

import "fmt"

// 切片
func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"中国"}
	fmt.Println(s1, s2)

	// 长度和容量
	fmt.Printf("len(s1): %d, cap(s1): %d\n", len(s1), cap(s1))

	// 由数组到切片，类似于 subList
	a1 := [...]int{1, 2, 3, 4, 5}
	s3 := a1[0:4]
	fmt.Println("s3:", s3)
	fmt.Printf("s3:%T\n", s3)
	s4 := a1[:4] // 等价于 a1[0:4]
	s5 := a1[3:] // 等价于 a1[3:len(a1)]
	s6 := a1[:]  // 等价于 a1[0:len(a1]
	fmt.Println(s4, s5, s6)
	// 切片可以动态扩容，底层指向的是数组
	fmt.Printf("len(s4):%d, cap(s4):%d; len(s5):%d, cap(s5):%d\n", len(s4), cap(s4), len(s5), cap(s5))
	s4 = append(s4, 5)
	fmt.Println(s4)

	// 切片再切片
	s7 := s4[1:3]
	fmt.Println("s7:", s7)
	a1[2] = 13000
	fmt.Println("s7:", s7) // 切片是一个引用类型，指向底层数组，改了底层数组，切片也会变

	// double cap when append ele
	slice0 := []string{"广州", "深圳", "上海"}
	fmt.Println("slice0 cap:", cap(slice0))
	slice0 = append(slice0, "北京") // 必须要用原始变量接收
	fmt.Println("slice0 cap after append:", cap(slice0))
	slice1 := []string{"成都"}
	fmt.Println("slice1 cap:", cap(slice1))
	slice1 = append(slice1, slice0...) // 扩容机制要了解
	fmt.Println("slice1 cap after append:", cap(slice1))

	// copy
	slice2 := make([]string, len(slice1), 10) // 长度也要写上
	size := copy(slice2, slice1)
	fmt.Println("size:", size)
	fmt.Println("slice1 after copying:", slice1)
	fmt.Println("slice2 after copying:", slice2)
}
