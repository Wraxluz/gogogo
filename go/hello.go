package main
import "fmt"
func main(){
	fmt.Println("HelloWorld")
	var s1 string
	s1 = "张三"
	fmt.Printf("%T,%s\n",s1,s1)
	s2 := `hello`
	fmt.Printf("%T,%s\n",s2,s2)
}