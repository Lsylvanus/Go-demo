package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	RandomTestBase()
}

//生成若干个不重复的随机数
func RandomTestBase() {
	//测试5次
	for i := 0; i < 5; i++ {
		nums := generateRandomNumber(10, 500, 10)
		time.Sleep(time.Second * 1)
		fmt.Println(nums)
	}
}
//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}