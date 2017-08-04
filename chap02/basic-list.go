package main

import (
	"fmt"
	"container/list"
)

func main() {
	// 生成队列
	l := list.New()

	// 入队, 压栈
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	// 出队
	i1 := l.Front()
	l.Remove(i1)
	fmt.Printf("%d\n", i1.Value)

	// 出栈
	i4 := l.Back()
	l.Remove(i4)
	fmt.Printf("%d\n", i1.Value)

	/*l := list.New()
	l.PushBack(path)

	var result []string
	for l.Len() > 0 {
		i1 := l.Front()
		l.Remove(i1)
		value, ok := i1.Value.(string)
		if !ok {
			fmt.Println("It's not ok for type string")
			continue
		}
		ret, e := d.SDK.Bucket(d.Bucket).Prefix(value).Delimiter(delimiter).HasCommonPrefix(true).ListObject()
		if e != nil {
			fmt.Printf("ListObject %v [%s]<%s> failed, %v", d.Bucket, prefix, delimiter, e)
		}

		for _, dir := range ret.CommonPrefixes {
			if strings.HasSuffix(dir, "/") {
				dir = strings.TrimSuffix(dir, "/")
			}
			result = append(result, dir)
			l.PushBack(dir+"/")
		}

		for _, file := range ret.Contents {
			result = append(result, file.Key)
		}
	}*/
}