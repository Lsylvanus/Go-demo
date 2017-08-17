package main

import (
	"fmt"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	combsort(items)
	fmt.Println(items)
}

/*
梳排序：
是冒泡排序算法的改进。
虽然冒泡排序总是比较相邻元素（gap = 1），
梳排序以 gap = n/1.3 开始，其中 n = len(items)，并在每次迭代时缩小 1.3 倍。

这种改进背后的想法是消除所谓的海龟（靠近列表末尾的小值）。
最后的迭代与 gap = 1 的简单冒泡排序相同。
*/
func combsort(items []int) {
	var (
		n       = len(items)
		gap     = len(items)
		shrink  = 1.3
		swapped = true
	)

	for swapped {
		swapped = false
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		for i := 0; i+gap < n; i++ {
			if items[i] > items[i+gap] {
				items[i+gap], items[i] = items[i], items[i+gap]
				swapped = true
			}
		}
	}
}