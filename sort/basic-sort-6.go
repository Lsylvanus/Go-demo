package main

import (
	"fmt"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	sortedItems := mergeSort(items)
	fmt.Println(sortedItems)
}

/*
归并排序：
是一种非常有效的通用排序算法。
这是分治算法的典型应用，这意味着列表被递归地分解成更小的列表，这些列表被排序然后被递归地组合以形成完整的列表。

来自维基百科：在概念上，合并排序的工作方式如下：
1.将未排序的列表划分为n个子列表，每个子列表包含1个元素（1个元素的列表被视为排序）。
2.重复合并子列表以生成新的排序子列表，直到只剩下1个子列表。 这将是排序列表。
*/
func mergeSort(items []int) []int {
	var n = len(items)

	if n == 1 {
		return items
	}

	middle := int(n / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, n-middle)
	)
	for i := 0; i < n; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// Either left or right may have elements left; consume them.
	// (Only one of the following loops will actually be entered.)
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
