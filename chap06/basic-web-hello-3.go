package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {

	type person struct {
		Id      int
		Name    string
		Country string
	}

	Lsyl := person{Id: 1001, Name: "Lsylvanus", Country: "China"}

	tmpl, err := template.ParseFiles("./userall.tpl", "./header.tpl", "./center.tpl", "./footer.tpl")
	if err != nil {
		fmt.Println("Error happened..")
	}
	err1 := tmpl.Execute(response, Lsyl)
	if err1 != nil {
		panic(err1.Error())
	}
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8874", nil))
}
