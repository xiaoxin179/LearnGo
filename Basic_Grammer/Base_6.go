package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数的使用
func oprator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		//panic就是中断执行
		panic("无法识别的操作符：" + op)

	}

}

// 返回两个参数的函数
func reTwo(a, b int) (int, int) {
	return a / b, a % b
}
func reTwo2(a, b int) (q, r int) {
	a = a / b
	r = a % b
	return

}

// 返回两个参数的第三种写法
func reTwo3(a, b int) (q, r int) {
	return a / b, a % b

}

// 函数式编程实现上述
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("call function %s with args"+"(%d,%d)", opName, a, b)
	return op(a, b)

}
func sum(number ...int) int {
	s := 0
	for i := range number {
		s += number[i]
	}
	return s

}

//可变参数列表

func main() {
	fmt.Println(oprator(12, 13, "*"))
	fmt.Println(reTwo(12, 5))
	q, r := reTwo2(12, 5)
	fmt.Println(q, r)
	fmt.Println(reTwo3(12, 5))
	fmt.Println("函数式编程：")
	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))

		}, 3, 4))
	fmt.Println("可变参数的列表：")
	fmt.Println(sum(12, 13, 13, 13, 23))

}
