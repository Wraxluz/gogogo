package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 4, 8)
	fmt.Println(s1)
	fmt.Printf("len =%d cap =%d slice = %v\n", len(s1), cap(s1), s1)

	var num []int

	if num == nil {
		fmt.Println("切片是空的")
	}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(num), cap(num), num)

	num = append(num, 0)
	num = append(num, 1, 2, 3, 4)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(num), cap(num), num)

	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("a      == ", a)
	fmt.Println("a[1:4] == ", a[1:4])
}
