package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

// 第二个条件结构
func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}

}

// 循环结构
func converTobin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

	}

}
func forerver() {
	for {
		fmt.Println("abc")

	}

}
func main() {
	fmt.Println(evel(12, 13, "*"))
	fmt.Println(grade(80))
	fmt.Println(bounded(123))
	fmt.Println("循环结构的使用:")
	fmt.Println(
		converTobin(1212),
		converTobin(10086),
		converTobin(4),
	)
	fmt.Println("打印文件内容：")
	//printFile("File.txt")
	forerver()

}
