package main

import "fmt"
import "bufio"


func main()  {
	/*
	输出：
		Pritn //打印
		Printf //格式化输出
		Println //打印后换行
	格式化打印占位符：
		%v，原样输出
		%T，打印字符串类型
		%t，bool类型
		%s，字符串
		%f，浮点
		%d，10进制
		%b，2进制
		%o，8进制
		%x，%X ， 16进制
		%c，打印字符
		%p， 打印地址
	输入：
		Scanln()
		Scanf()
	文件读写：
		bufio包
	*/

	/*
	var x int
	var y float64
	fmt.Println("输入一个整数和一个浮点数：")
	fmt.Scanln(&x,&y)//必须要有&符号才能使用 Scanln输入多个数值时用空格隔开
	fmt.Printf("%d,%f\n",x,y)
	fmt.Scanf("%d,%f\n",&x,&y)//Scanf输入多个数值需要用逗号隔开
	fmt.Printf("%d,%f\n",x,y)
	*/
	reader:= bufio.NewReader(os.Stdin)
	s1,_:=reader.ReadString('\n')
	fmt.Println(s1)
}