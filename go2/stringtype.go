package main
import "fmt"
func main(){
	/*
	如果想输出 "hello" 需要用转义字符:\   如：fmt.Println("\"hello\"")

	\t是加空格，\n是换行
	*/
	var s1 string
	s1 = "张三"
	fmt.Printf("%T,%s\n",s1,s1)
	s2 := `hello`
	fmt.Printf("%T,%s\n",s2,s2)
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%d\n",v1,v1)
	fmt.Printf("%T,%s\n",v2,v2)
	
	v3 := '家'
	fmt.Printf("%T,%d.%c,%q\n",v3,v3,v3,v3)	
	fmt.Println("\"hello\"")
	fmt.Println("Hello\nWorld")
	fmt.Println("Hello\tWorld")
}