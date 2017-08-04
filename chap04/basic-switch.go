package main

import "fmt"

func main() {
	//pattarn 1
	weekday := "FRI"
	switch weekday {
	case "MON":
		fmt.Println("Monday...")
	case "TUE":
		fmt.Println("Tuesday...")
	case "WED":
		fmt.Println("Wednesday...")
	case "THU":
		fmt.Println("Thursday...")
	case "FRI":
		fmt.Println("Friday...")
	case "SAT":
		fmt.Println("Saterday...")
	case "SUN":
		fmt.Println("Sunday...")
	default:
		fmt.Println("Default branche...")
	}

	//pattarn 2
	switch weekday := "SUN"; weekday {
	case "MON":
		fmt.Println("Monday...")
	case "TUE":
		fmt.Println("Tuesday...")
	case "WED":
		fmt.Println("Wednesday...")
	case "THU":
		fmt.Println("Thursday...")
	case "FRI":
		fmt.Println("Friday...")
	case "SAT":
		fmt.Println("Saterday...")
	case "SUN":
		fmt.Println("Sunday...")
	default:
		fmt.Println("Default branche...")
	}

	//pattarn 3
	score := 75
	switch {
	case score < 60:
		fmt.Println("Grade D...")
	case score <= 70:
		fmt.Println("Grade C...")
	case score <= 80:
		fmt.Println("Grade B...")
	case score <= 90:
		fmt.Println("Grade A...")
	default:
		fmt.Println("Default branche...")
	}

	//pattarn 4
	var obj interface{}
	obj = "Hello World"
	switch obj.(type) {
	case int:
		fmt.Println("This is a int type ...")
	case string:
		fmt.Println("This is a string type...")
	default:
		fmt.Println("Default branch...")
	}
}
