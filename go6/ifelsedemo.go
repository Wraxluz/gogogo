package main

import "fmt"

func main() {
	/*
		score := 0
		fmt.Println("请输入成绩:")
		fmt.Scanln(&score)

			if score >= 80 {
				fmt.Println("优秀")
			} else {
				if score >= 60 {
					fmt.Println("及格")
				} else {
					fmt.Println("不及格")
				}
			}
	*/
	if a := 61; a >= 80 {
		fmt.Println("优秀")
	} else if a >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
	num := 2
	switch num {
	case 1:
		fmt.Println("优秀")
	case 2:
		fmt.Println("及格")
		fallthrough
	case 3:
		fmt.Println("不及格")
	default:
		fmt.Println("输入错误")
	}
	fmt.Println("————————————————")
	letter := "B"
	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println("是元音")
	default:
		fmt.Println("不是元音")
	}
}
