package main

import "fmt"

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
func main() {
	fmt.Println(oprator(12, 13, "*"))
	fmt.Println(reTwo(12, 5))
	q, r := reTwo2(12, 5)
	fmt.Println(q, r)
	fmt.Println(reTwo3(12, 5))

}
