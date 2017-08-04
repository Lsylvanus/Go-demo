package main

import "fmt"

func main() {
	//bool var definition and init
	var flag bool = false
	fmt.Println("flag =", flag)

	//int var
	var i, j, k int
	i = 1
	j = 2
	k = 3
	fmt.Printf("i=%d, j=%d, k=%d\n", i, j, k)

	//var rune and init
	var r1, r2, r3 rune = 1, 2, 3
	r4, r5, r6 := 4, 5, 6
	fmt.Printf("r1=%d r3=%d r3=%d\n", r1, r2, r3)
	fmt.Printf("r4=%d r5=%d r6=%d\n", r4, r5, r6)

	//var float32 and init
	var (
		f1 float32 = 2.1
		f2 float32 = 3.4
	)
	fmt.Printf("f1=%f f2=%f\n", f1, f2)

	//var string and init
	var str1 string
	str1 = "Lsylvanus"
	fmt.Println("hello,", str1)

	//:=
	str2 := "hello"
	fmt.Printf("str2 = %s \n", str2)
}