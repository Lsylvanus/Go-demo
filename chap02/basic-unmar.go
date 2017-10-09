package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var dat map[string]interface{}
	data := fmt.Sprintf("%v", "{{\"sku\":\"22222\"，\"sum\":\"50\"，\"conut\":\"100\"，\"purchaseagent\":\"谭\"}，{\"sku\":\"33333\"，\"sum\":\"60\"，\"count\":\"200\"，\"purchaseagent\":\"张\"}}")
	err := json.Unmarshal([]byte(data), &dat)
	if err != nil {
		fmt.Println("err :", err)
	}
}