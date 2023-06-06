package main

import "fmt"

// go语言的指针
// 指针他不参与运算
// 参数传递。在go语言中只有值传递一种机制，掉函数的时候参数都是需要拷贝一份的
func swap(a, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}
func main() {
	a := 12
	var pa *int = &a
	*pa = 1
	fmt.Println(a)
	b := 45
	swap(&a, &b)
}
