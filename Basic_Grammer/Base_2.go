package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 内建函数类型bool,string，(u)int ,(u)int16,uintptr(自增)
// byte,rune:相当于其他语言中的char
// float32, float64, complex64, complex128：复数，实部虚部分别占据64位
// i等于根号-1
func eulter() {
	c := 3 + 4i
	//cmplx是一个内置复数数学库
	//cmplx.Abs(): 计算复数的模（即绝对值）
	//cmplx.Conj(): 计算复数的共轭
	//cmplx.Exp(): 计算复数的指数函数
	//cmplx.Log(): 计算复数的自然对数
	//cmplx.Sqrt(): 计算复数的平方根
	//cmplx.Phase(): 计算复数的辐角
	fmt.Println(cmplx.Abs(c))
}

// 在其中使用欧拉公式
func enter() {
	fmt.Println(cmplx.Exp((1i * math.Pi) + 1))
}

// 强制类型转换，没有隐式类型转换
func trangle() {
	var a, b int = 1, 2
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}
func main() {
	trangle()
}
