package main

import "fmt"

// 数组
func main() {
	var a1 [2]bool
	fmt.Printf("%T \n", a1)

	// 定长初始化
	a2 := [2]bool{true, false}
	fmt.Println(a2)

	// 不定长
	a3 := [...]bool{false, true, false, true}
	fmt.Println(a3)
	a3[0] = true
	// 遍历数组
	for i, b := range a3 {
		fmt.Println(i, b)
	}

	// 二维数组
	aa1 := [...][2]int{
		{3, 4},
		{5, 6},
	}
	for i, ints := range aa1 {
		fmt.Println(i, ints)
	}

	// 题目
	// 求和
	a4 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, e := range a4 {
		sum += e
	}
	fmt.Println("sum:", sum)

	// 寻找两数之为 8 的索引对
	for i, v := range a4 {
		ans := 8 - v
		for j := i + 1; j < len(a4); j++ {
			if a4[j] == ans {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}

}
