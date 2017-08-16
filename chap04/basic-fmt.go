package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "ysn0420007_S_Yellow"
	matched, err := regexp.MatchString("(.*)([^\\[]*)\\](.*)", str)
	fmt.Println(matched, err)
	fmt.Println(strings.HasSuffix("NLT_abc", "abc")) //后缀是以NLT开头的
	fmt.Println(strings.Index("NLT_abc", "_")) // 返回第一个匹配字符的位置，这里是3

	fmt.Printf("%q\n", strings.SplitN("Y[N0260854]FA", "[", -1)) //
	str1 := strings.SplitN("Y[N0260854]FA", "[", -1)
	var str2 string
	for k, v := range str1 {
		if strings.Contains(v, "]") {
			fmt.Println("contain",k)
			str2 = v
		}else {
			fmt.Println(k)
		}
	}
	str3 := strings.SplitN(str2, "]", -1)
	fmt.Println(str3)
	for k, v := range str3 {
		if k == 0 {
			str2 = v
		}
	}
	fmt.Println(str2)
}