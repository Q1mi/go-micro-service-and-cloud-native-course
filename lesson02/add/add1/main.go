package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	x := 10
	y := 20
	ret := add(x, y) // 本地调用
	fmt.Println(ret)
}
