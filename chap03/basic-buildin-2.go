package main

import "fmt"

func main() {
	a1 := new([10]int)
	a1[4] = 5
	a1[7] = 8
	fmt.Println("a1= ", a1, "len(a1)=", len(a1), " cap(a1)=", cap(a1))
	fmt.Println("------------------------------------------")
	s1 := make([]int, 5, 10)
	s1[0] = 5
	s1[4] = 2
	s2 := make([]int, 5, 10)
	s2[0] = 1
	s2[4] = 3
	fmt.Println("before copy :s1= ", s1, "len(s1)=", len(s1), " cap(s1)=", cap(s1))
	fmt.Println("before copy :s2= ", s2, "len(s2)=", len(s2), " cap(s2)=", cap(s2))
	fmt.Println("------------------------------------------")
	println("copy(s2, s1)")
	copy(s2, s1)
	fmt.Println("after  copy :s1= ", s1, "len(s1)=", len(s1), " cap(s1)=", cap(s1))
	fmt.Println("after  copy :s2= ", s2, "len(s2)=", len(s2), " cap(s2)=", cap(s2))
	fmt.Println("------------------------------------------")
	println("reset")
	s2[0] = 1
	s2[4] = 3
	fmt.Println("after  reset:s1= ", s1, "len(s1)=", len(s1), " cap(s1)=", cap(s1))
	fmt.Println("after  reset :s2= ", s2, "len(s2)=", len(s2), " cap(s2)=", cap(s2))
	fmt.Println("------------------------------------------")
	println("append(s2, 20)")
	s2 = append(s2, 6)
	fmt.Println("after  append:s1= ", s1, "len(s1)=", len(s1), " cap(s1)=", cap(s1))
	fmt.Println("after  append:s2= ", s2, "len(s2)=", len(s2), " cap(s2)=", cap(s2))
	fmt.Println("------------------------------------------")
	s2 = append(s2, 7)
	s2 = append(s2, 8)
	s2 = append(s2, 9)
	s2 = append(s2, 10)
	fmt.Println("after  append:s2= ", s2, "len(s2)=", len(s2), " cap(s2)=", cap(s2))
	s2 = append(s2, 11)
	fmt.Println("after  append:s2= ", s2, "len(s2)=", len(s2), " cap(s2)=", cap(s2))

}
