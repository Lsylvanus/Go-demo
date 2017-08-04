package main
func main() {

	//pattarn 1
	score := 85
	if score < 60 {
		println("Grade1: D")
	} else if score < 70 {
		println("Grade1: C")
	} else if score < 80 {
		println("Grade1: B")
	} else if score <= 100 {
		println("Grade1: A")
	} else {
		println("Grade: undefined")
	}

	//pattarn 2
	if score := 75; score < 60 {
		println("Grade2: D")
	} else if score < 70 {
		println("Grade2: C")
	} else if score < 80 {
		println("Grade2: B")
	} else if score <= 100 {
		println("Grade2: A")
	} else {
		println("Grade: undefined")
	}
}
