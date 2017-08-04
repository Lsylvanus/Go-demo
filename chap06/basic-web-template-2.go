package main

import (
	"fmt"
	"os"
	"html/template"
)

func main() {

	type person struct {
		Id      int
		Name    string
		Country string
	}

	Lsyl := person{Id: 1001, Name: "Lsylvanus", Country: "China"}

	fmt.Println("Lsyl = ", Lsyl)

	tmpl, err := template.ParseFiles("./tmpl1.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(os.Stdout, Lsyl)

}