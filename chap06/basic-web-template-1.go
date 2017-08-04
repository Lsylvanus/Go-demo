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

	var p person
	p.Id = 123
	p.Name = "Tom"
	p.Country = "guangdong"

	p1 := person{Id: 1001, Name: "Lsylvanus", Country: "China"}

	fmt.Println("Tom = ", p)
	fmt.Println("Lsyl = ", p1)

	tmpl := template.New("lsyl")
	/*for i := 1; i <= p.len; i++ {
		fmt.Println(p.Name)
		tmpl.Execute(os.Stdout, p.Name)
		tmpl.Parse("Hello {{.Name}} Welcome to go programming...\n")
	}*/
	tmpl.Parse("Hello {{.Name}} Welcome to go programming...\n")
	tmpl.Execute(os.Stdout, p)
}