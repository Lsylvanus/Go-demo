package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func Hello(response http.ResponseWriter, request *http.Request) {

	type person struct {
		Id      int
		Name    string
		Country string
	}

	Lsyl := person{Id: 1001, Name: "Lsylvanus", Country: "China"}

	tmpl, err := template.ParseFiles("./userall.tpl","./header.tpl","./center.tpl","./footer.tpl")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(response, Lsyl)
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}