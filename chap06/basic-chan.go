package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//ThreadItem 帖子数据
type ThreadItem struct {
	url     string   //帖子路径
	content string   //帖子内容
	imgs    []string //帖子图片
}

var (
	//图片正则表达式
	imageItemExp = regexp.MustCompile(`src="//i\.4cdn\.org/s/[0123456789]+s\.jpg"`)
	//帖子路径正则表达式
	threadItemExp = regexp.MustCompile(`"thread/[0123456789]+"`)
)

func httpGet(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

func get(content string) []ThreadItem {
	//解析正则表达式，返回参数tds是一个数组
	tds := threadItemExp.FindAllStringSubmatch(content, 10000)
	var tdStr = make([]string, 0)
	//去掉引号，并放到一维数组中
	for _, t := range tds {
		var n = strings.Replace(t[0], "\"", "", -1)
		tdStr = append(tdStr, n)
	}
	var threads = make([]ThreadItem, 0)
	//组装帖子结构体
	for _, t := range tdStr {
		threads = append(threads, ThreadItem{url: "http://boards.4chan.org/s/" + t})
	}
	return threads
}
