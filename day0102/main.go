package main

import (
	"fmt"
	"time"
)

// 全局变量
var (
	name string
	age  int
)

// 批量声明常量
const (
	OK       = 200
	NOTFOUND = 404
	ERROR    = 500
	ERROR2          // 没有赋值表示跟上面一样
	_        = iota // 匿名变量
	A1
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func main() {
	name = "fuxin"
	age = 10
	fmt.Printf("Hello, %s! Are you %d years old?\n", name, age)

	// 类型推断
	var s0 = "推断"
	fmt.Println(s0)

	// 简短变量
	s1 := "test"
	fmt.Println(s1)
	s1 = "变量"

	fmt.Println(s1)

	duration := time.Second
	fmt.Println(duration)

	// 常量
	const pi = 3.14159
	fmt.Println(pi)
	fmt.Println("常量:", OK)
	fmt.Println("常量未赋值:", ERROR2)

	// iota
	fmt.Println("iota:", A1) // iota 对每一行常量声明 + 1
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("TB:", TB)

	// 整型
	var num int8 = 127 // byte
	fmt.Println(num)
	var num2 int16 = 126 // short
	fmt.Println("num2:", num2, "len:")
	num3 := 077 // int
	var num4 int64 = 234
	fmt.Println(num3)
	fmt.Printf("%T\t%T\t%T\t%T\n", num4, num3, num2, num)

	// fmt 占位符
	var n = 100
	fmt.Println("------------- fmt 占位符 ---------------")
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%d\n", n)

}
