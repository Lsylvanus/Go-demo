package main

import (
	"fmt"
)

func main() {
	items := []int{4, 202, 3, 9, 6, 5, 1, 43, 506, 2, 0, 8, 7, 100, 25, 4, 5, 97, 1000, 27}
	shellshort(items)
	fmt.Println(items)
}

/*
希尔排序：
首先，选择一连串的间隙。
有许多不同的公式来产生间隙序列，算法的平均时间复杂度取决于这个变量。
例如，我们选择 (2 ^ k ) – 1，前缀为1，这将给我们[1,3,7,15,31,63 ...]。
反转顺序：[...，63，31，15，7，3，q]。

现在遍历颠倒的间隙列表，并在每个子列表中使用插入排序。
所以在第一次迭代中，每第63个元素都应用插入排序。
在第二次迭代中，每31个元素应用插入排序。
所以一路下来到1。最后一次迭代将在整个列表中运行插入。
*/
func shellshort(items []int) {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := pow(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}
	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap] > items[j] {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}