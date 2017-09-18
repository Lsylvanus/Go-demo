package main

import (
	"ERP_1/log"
	"flag"
	"time"

	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"goTest/db"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	page     = 50
	sort     = "normal"
	acc_key  = "b637d49acf7c9644fc7d39d11e894fee"
)

func getApiKey(engine *xorm.Engine) string {
	key := new(db.ApiUser)
	_, err := engine.Where("user_code =?", "289").Get(key)
	if err != nil {
		log.Log.Error("get api key :", err)
		return acc_key
	}
	return key.AccessKey
}

func getResp(engine *xorm.Engine, examType string, pagenum, subject int) []byte {
	urlStr := "https://way.jd.com/jisuapi/driverexamQuery?type=" + examType + "&subject=" + strconv.Itoa(subject) + "&pagesize=" + strconv.Itoa(page) + "&pagenum=" + strconv.Itoa(pagenum) + "&sort=" + sort + "&appkey=" + getApiKey(engine)
	req, _ := http.NewRequest("GET", urlStr, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Log.Error("client do req :", err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func Unmalshal(engine *xorm.Engine) {
	submap := make(map[int]int)
	submap[0] = 1
	submap[1] = 4

	examTypes := make(map[int]string)
	examTypes[0] = "a1"
	examTypes[1] = "a2"
	examTypes[2] = "a3"
	examTypes[3] = "b1"
	examTypes[4] = "b2"
	examTypes[5] = "c2"
	examTypes[6] = "c3"
	examTypes[7] = "d"
	examTypes[8] = "e"
	examTypes[9] = "f"

	for _, subject := range submap {
		pagenum := 1
		for _, examType := range examTypes {
			for {
				driver := new(db.DriverEx)
				body := getResp(engine, examType, pagenum, subject)
				err := json.Unmarshal(body, driver)
				if err != nil {
					log.Log.Error("json unmarshal :", err)
					return
				}

				//total := driver.Result.ExamRe.Total
				examList := driver.Result.ExamRe.List
				for _, exam := range examList {
					in := new(db.ExamList)
					if in.Type != "" {
						in.Type = exam.Type
					} else {
						in.Type = examType
					}
					in.Answer = exam.Answer
					in.Pic = exam.Pic
					in.Chapter = exam.Chapter
					in.Explain = exam.Explain
					in.Question = exam.Question
					in.Option1 = exam.Option1
					in.Option2 = exam.Option2
					in.Option3 = exam.Option3
					in.Option4 = exam.Option4
					in.Subject = strconv.Itoa(subject)

					_, err := engine.Insert(in)
					if err != nil {
						log.Log.Error("insert into exam_list :", err)
						return
					}
				}
				pagenum++
				time.Sleep(time.Second * 10)
				if len(examList) < page {
					break
				}
			}
		}
	}

}

func main() {
	log.LogInit()
	dbConn := flag.String("db", "root:chenghai3c@/express_delivery?charset=utf8&collation=utf8_general_ci&parseTime=true", "MySQL DB path") //zhusq
	flag.Parse()
	// 连接数据库
	var err error
	engine, err := xorm.NewEngine("mysql", *dbConn)
	if err != nil {
		log.Log.Panic(err)
	}
	err = engine.Ping()
	if err != nil {
		log.Log.Panic(err)
	}
	log.Log.Info("success to connect db")

	cst, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Log.Error(err)
	}
	engine.DatabaseTZ = cst //设置数据库的系统时区
	time.Local = cst        //设置time的时区

	Unmalshal(engine)
}
