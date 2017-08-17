package main

import (
	"fmt"
	"time"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	bubbleSort(items)
	fmt.Println(items)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

/*
冒泡排序：
在循环中，从第一个元素到第 n 个（n = len(items)）迭代数组。
比较相邻的值，如果它们的顺序错误，交换它们。
您可以通过在每次迭代后将 n 递减 1 来优化算法。
*/
func bubbleSort(items []int) {
	var (
		n       = len(items)
		swapped = true
	)
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		n = n - 1
	}
}