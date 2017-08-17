package main

import (
	"fmt"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	insertionSort(items)
	fmt.Println(items)
}

/*
插入排序：
* 自适应：时间复杂度随着已经基本排序的列表而减少 – 如果每个元素不超过其最终排序位置的 k 个位置，则 O(nk)
* 稳定：相等值的索引的相对位置不变
* 就地：只需要一个常数 O(1) 的额外的内存空间
* 在实践中比泡沫或选择排序更有效
*/
func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}