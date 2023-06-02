package main

//go语言定义变量的各种方法
import "fmt"

// 包内变量，在go语言中不存在全局变量
var sss = 3
var (
	b1 = 12
	b2 = 13
	b3 = "sdfdsdf"
)

func mul() {
	//定义的变量是一定要使用到的
	var a, b int = 1, 2
	var s string = "dadas"
	fmt.Println(a, b)
	fmt.Println(s)

}
func mul2() {
	//不规定类型也是可以写在一行
	var a = 1
	var b = "fdfsf"
	fmt.Println(a, b)

}
func mul3() {
	//	四年一冒号的时候可以直接使用：来简化操作
	//冒号等于只可以在函数中使用，不可以在包内使用

	a, b, c, d := 3, 4, "dfsfsf", 10
	fmt.Println(a, b, c, d)

}
func main() {
	mul()
	mul2()
	fmt.Println("mul3函数：")
	mul3()
	fmt.Println("ss=", sss)
	fmt.Println("对包内变量的使用：")
	fmt.Println(b1, b2, b3)
}
