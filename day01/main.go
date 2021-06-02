package main

import (
	"fmt"
	"time"
)

func main() {
	// 导入
	fmt.Println("Hello world!")
	fmt.Println("东霖是zz")
	fmt.Printf("test-%s-%d", "llll", 100)
	fmt.Println()
	duration := 2 * time.Second
	time.Sleep(duration)
	var test = duration
	fmt.Println(test.String())
}
