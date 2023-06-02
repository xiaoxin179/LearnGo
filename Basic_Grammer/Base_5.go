package main

import "fmt"

func evel(a, b int, op string) int {
	var result int
	//不需要贾break，人性化
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		//panic是抛出异常的函数
		panic("无法识别的运算符号：" + op)

	}
	return result

}
func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "B"
	case score < 100:
		g = "A"
	default:
		panic("不要随便输入")
	}
	return g

}
//循环结构
func
func main() {
	fmt.Println(evel(12, 13, "*"))
	fmt.Println(grade(80))
}
