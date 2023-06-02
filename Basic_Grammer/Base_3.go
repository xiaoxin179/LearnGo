package main

import (
	"fmt"
	"math"
)

//常量定义

func consts() {
	//const定义的变量在这里都是可以直接当作文本来处理的
	const filename = "abc.txt"
	//不赋值类型
	fmt.Println("常量a，b:")
	const a, b = 3, 4
	//本来sqrt函数是需要类型转换的，但是在这里的时候本来
	fmt.Println(math.Sqrt(a + b))
	//	直接定义c,d
	fmt.Println("直接定义c,d:")
	c, d := 3, 4
	fmt.Println(math.Sqrt(float64(c + d)))

}

// 一组常量的定义
func consts2() {
	const (
		yuanzhou = 3.14
		a, b     = "a", "b"
	)
	fmt.Println(yuanzhou)
}

// 枚举类型通过const来定义
func enums() {
	const (
		//自动增加
		cpp = iota
		java
		golang
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		tb
		pb
	)
	fmt.Println(cpp, java, golang)
	fmt.Println(b, kb, mb, tb, pb)

}
func main() {
	consts()
	consts2()
	enums()
}
