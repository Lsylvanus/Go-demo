package main
func main() {
	//pattarn 1
	for level := 0; level < 10; level++ {
		for cnt := 0; cnt < 10-level; cnt++ {
			print(" ")
		}
		for ascnt := 0; ascnt < level; ascnt++ {
			print("*")
		}
		println("")
	}

	//pattarn 2
	level := 0
	for level < 10 {
		for cnt := 0; cnt < 10 - level; cnt++ {
			print(" ")
		}
		for ascnt := 0; ascnt < 2 * level - 1; ascnt++ {
			print("*")
		}
		println("")
		level++
	}

	//pattarn 3
	/*
			for {
					println("please don try this one unless you want to test your cpu usage")
			}
	*/

}
