package main

import (
	//"Go-demo/win/jwt"
	//"github.com/Sirupsen/logrus"
	"fmt"
	"github.com/ajstarks/svgo"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//db := jwt.OpenSql()
	//jwt.FindDb(db)
	//db.Close()
	//jwt.PrintLs()

	//for i := 0; i < 5; i++ {
	//	c := comm{Count:10}
	//	str := c.getRandNum()
	//	fmt.Println("str :", str)
	//}

	/*width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	canvas.Ellipse(width / 2, height, width / 2, height / 3, "fill:rgb(44,77,232)")
	canvas.Text(width / 2, height / 2, "Hello，World", "fill：white; font-size：48pt; text-anchor：middle")
	canvas.End()*/

	width := 200
	height := 200
	w := 1.0
	wi := 0.03
	vi := 15.0
	h2 := height / 2
	w2 := width / 2

	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	canvas.Gstyle("font-family:serif; fill:white; font-size:72pt")
	canvas.Gtransform(fmt.Sprintf("translate(%d, %d)", w2, h2))

	for i := 0.0; i <= 360.0; i += vi {
		canvas.Text(0, 0, "i",
			fmt.Sprintf(`transform ="rotate(%.3f)"`, i),
			fmt.Sprintf(`fill-opacity="%.3f"`, w))
		w -= wi
	}
	canvas.Gend()
	canvas.Gend()
	canvas.End()
}

type comm struct {
	Count int
}

func (c comm) getRandNum() string {
	nums := make([]string, 0)
	// 随机数种子
	for len(nums) < c.Count {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//生成随机数
		num := r.Intn(100)
		n := strconv.Itoa(num)
		//查重
		exist := false
		for _, v := range nums {
			if v == n {
				exist = true
				break
			}
		}
		// 不存在append
		if !exist {
			nums = append(nums, n)
		}
	}
	return strings.Join(nums, "")
}