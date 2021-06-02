package main

import "fmt"

/*
  流程控制
*/
func main() {
	age := 10
	// if 语句

	if age >= 20 {
		fmt.Println("我20岁了")
	} else if age >= 10 {
		fmt.Println("我10岁了")
	} else {
		fmt.Println("我刚出生")
	}

	// for 语句
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}

	// for range
	s := "hello go 语言"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// switch
	switch age {
	case 10:
		fmt.Println("我10岁了")
	case 20:
		fmt.Println("我20岁了")
	default:
		fmt.Println("我刚出生")
	}

}
