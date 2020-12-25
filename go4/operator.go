package main

import "fmt"

func main()  {
//算数运算符 +,-,*,/,%,++,--
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d\n",a,b,sum)

	sub := a- b
	fmt.Printf("%d - %d = %d\n",a,b,sub)

	mul := a*b
	fmt.Printf("%d * %d = %d\n",a,b,mul)

	div := a / b
	fmt.Printf("%d / %d = %d\n",a,b,div)

	mod := a % b
	fmt.Printf("%d %% %d = %d\n",a,b,mod)
//关系运算符 ==,>,<,<=,>=,!=
	v1 := 3
	v2 := 5
	v3 := 3
	res1 := v1 > v2
	res2 := v2 > v3
	fmt.Printf("%T,%t\n",res1,res2)

//逻辑运算符 &&(与),||(或),!(非)
	f1 := true
	f2 := false
	f3 := true
	ra1 := f1&&f2
	ra2 := f1&&f3
	fmt.Println(ra1,ra2)
	fmt.Println(f1 || f2)

//位运算符 按位与:& , 按位或: | , 异或: ^ ^a按位取反: 1——>0

	a1 := 60
	a2 := 13
	fmt.Printf("%d,%b\n",a1,a1)
	fmt.Printf("%d,%b\n",a2,a2)
	fmt.Println(a1 & a2)
	fmt.Println(a1 | a2)
	fmt.Println(a1 ^ a2)
	fmt.Println(a1 >> 2)
	fmt.Println(a1 << 2)
}