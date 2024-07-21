package main

import (
	"fmt"
	"leetcode-go/leetcode/doublePointer" // 显式导入包
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := doublePointer.Trap(height)
	fmt.Println(result)
}
