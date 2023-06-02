package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "File.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Println(string(contents))
	} else {
		fmt.Println("can not print file contents:", err)
	}

}
