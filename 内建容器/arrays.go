package main

import "fmt"

// 数组切片和容器得使用
// 数组的遍历
func bianliArr() {
	numbers := [6]int{12, 121, 121, 13, 232, 23}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])

	}

}
func printArr(arr [3]int) {
	for i, v := range arr {
		fmt.Println(i, v)

	}

}
func main() {
	bianliArr()
	arr := [3]int{12, 12, 34}
	printArr(arr)

}
