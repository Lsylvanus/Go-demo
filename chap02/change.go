package main

import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	"strconv"
	"regexp"
)

func Max(a, b int) bool {
	for {
		if a>b {
			return true
		}
		return false
	}
}

func SqlitComma(s, seq string) {
	strs := strings.Split(s, seq)
	for k, v := range strs {
		fmt.Printf("key is :%v,\t value is %v\n", k, v)
	}
	fmt.Println("len is :", len(strs))
	a := ""
	for i := 1; i <= len(strs); i++ {
		if strs[i-1] != "" {
			a += "&AmazonOrderId.Id."+ strconv.Itoa(i) + "=" + strs[i-1]
		}
	}
	fmt.Println("a is :", a)
}

func SplitEqual(s, seq string) (string, string) {
	var forward, behind string
	str := strings.Split(s, seq)
	for i, v := range str {
		if i == 0 {
			forward = v
		} else {
			behind = v
		}
	}
	return forward, behind
}

func SplitWord(key string) {
	str := strings.Split(key, "=")
	for i, s := range str {
		fmt.Printf("i is :%v, s is :%v\n", i, s)
	}
}

func main() {
	key := "MarketplaceId.Id.1=rerr342dgdh"
	SplitWord(key)
	f, b := SplitEqual(key, "=")
	fmt.Printf("f is :%v, b is :%v\n", f, b)

	keys := "JYQHSFK7GCCK5BWI4OE62KHG6MQHIB3KGHDA"
	ks := strings.Split(keys, ";")
	for i, v := range ks {
		fmt.Printf("i is :%v, v is :%v\n", i, v)
	}
	
	fmt.Println(Max(4, 3))

	/*old := "d:\\wish-get.xml"
	news := "d:\\wish-get.xmlcc"
	RenameFile(old, news)*/

	fmt.Println(getCurrentPath())
	
	orderIdStr := "113-7709084-1983410,113-0364739-2301803,113-9299225-8234620,114-5401575-7739461,113-7079822-8737021,112-7376127-7528262,111-6849940-8256250,114-3751159-0655422,113-5446059-2244227,114-2533908-8661065,111-4656205-6164209,"
	SqlitComma(orderIdStr, ",")

	regex := "^([hH][tT]{2}[pP]:/*|[hH][tT]{2}[pP][sS]:/*|[fF][tT][pP]:/*)(([A-Za-z0-9-~]+).)+([A-Za-z0-9-~\\/])+(\\?{0,1}(([A-Za-z0-9-~]+\\={0,1})([A-Za-z0-9-~]*)\\&{0,1})*)$"
	str := "https://developers.linode.com/list/serverInfo?s.11=vultr&id.11=rergfg&s.12=linode&id.12=gfgrtr"
	if getRegex(regex, str) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	regex1 := "(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]"
	fmt.Println(getRegex(regex1, str))

	/*c := 0
	for {
		if getRegex(regex, str) {
			c++
			fmt.Println(c)
			continue
		}
		c++
		fmt.Println("c is:", c)
		if c == 10 {
			break
		}
	}*/

	for i := 0; i <= 10 ; i++ {
		if getRegex(regex, str) {
			fmt.Println(i)
			continue
		}
		fmt.Println("i is:", i)
	}

	for i := 0; i <= 10 ; i++ {
		if i != 11 {
			continue
			fmt.Println("continue ", i)
		}
	}
	fmt.Println("finished.")
}

func RenameFile(oldName, newName string) error {
	err := os.Rename(oldName, newName)
	if err != nil {
		fmt.Println("file rename err :", err)
	}
	return nil
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println("path :", err)
		return ""
	}
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func getRegex(regex, str string) bool {
	b, err := regexp.MatchString(regex, str)
	if err != nil {
		return false
	}
	if b {
		return b
	}
	return  false
}