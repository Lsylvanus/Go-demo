package main

import (
	"fmt"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	selectionSort(items)
	fmt.Println(items)
}

/*
选择排序：
选择排序可以通过两个嵌套 for 循环来实现。
外部循环遍历列表 n 次（n = len(items)）。
内部循环将始终以外部循环的当前迭代器值开始（因此在每个迭代中，它将从列表中的更右侧的位置开始），
并找出子列表的最小值。
使用找到的最小值交换子列表的第一项。
*/
func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}