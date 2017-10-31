package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	/*var dat map[string]interface{}
	data := fmt.Sprintf("%v", "{{\"sku\":\"22222\"，\"sum\":\"50\"，\"conut\":\"100\"，\"purchaseagent\":\"谭\"}，{\"sku\":\"33333\"，\"sum\":\"60\"，\"count\":\"200\"，\"purchaseagent\":\"张\"}}")
	err := json.Unmarshal([]byte(data), &dat)
	if err != nil {
		fmt.Println("err :", err)
	}*/

	var keys []string
	var d string = "[\"JYQHSFK7GCCK5BWI4OE62KHG6MQHIB3KGHDA\",\"M4YHTJKCWV3NXGOIK7UGARN2K7VYXZJJGMUQ\"]"
	err1 := json.Unmarshal([]byte(d), &keys)
	if err1 != nil {
		fmt.Println("err1 :", err1)
	}
	fmt.Println("keys :", keys)

	nums := strings.Split("2255533;225111", ",")
	var param string
	for _, num := range nums {
		fmt.Println("num :", num)
		param += "&ids%5B%5D=" + num
	}
	fmt.Println("param :", param)
}