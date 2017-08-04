package main

import "net/http"

func Hello(resp http.ResponseWriter, req *http.Request)  {
	resp.Write([]byte("<b>Hello, Welcome to go web programming...</b><br>"))
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe("localhost:8888", nil)
}