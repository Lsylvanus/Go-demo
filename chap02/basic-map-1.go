package main

import (
	"fmt"
	"strconv"
)

func ranMap(reMap map[int][]string) {
	var name, id string
	for _, pMap := range reMap {
		for k, v := range pMap {
			if k%2 == 1 {
				id = v
				fmt.Println("id :", id)
				fmt.Println("name :", name)
			} else {
				name = v
			}
		}
	}
}

func main() {
	var param []string
	paramMap := make(map[int][]string)
	ch := 1
	for {
		s := "s." + strconv.Itoa(ch)
		id := "id." + strconv.Itoa(ch)
		if ch > 2 || (s == "" && id == "") {
			break
		}
		if s != "" && id != "" {
			param = append(param, s, id)
			fmt.Println("param :", param)
		}
		ch++
	}
	paramMap[ch] = param
	ranMap(paramMap)
}