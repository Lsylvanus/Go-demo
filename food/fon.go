package main

import (
	"ERP_1/log"
	"encoding/json"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"goTest/db"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	acc_key = "b637d49acf7c9644fc7d39d11e894fee"
	num     = 20
)

func Unmal(engine *xorm.Engine) {
	//inParentChild(engine)
	typesSearch(engine)
}

func getApiKey(engine *xorm.Engine) string {
	key := new(db.ApiUser)
	_, err := engine.Where("user_code =?", "289").Get(key)
	if err != nil {
		log.Log.Error("get api key :", err)
		return acc_key
	}
	return key.AccessKey
}

func getRespBody(url string) []byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func inParentChild(engine *xorm.Engine) {
	url := "https://way.jd.com/jisuapi/recipe_class?appkey=" + getApiKey(engine)
	data := getRespBody(url)
	bcc := new(db.BccData)
	err := json.Unmarshal(data, bcc)
	if err != nil {
		log.Log.Error("unmarshal json :", err)
		return
	}

	parents := bcc.Result.ReS
	for _, p := range parents {
		parent := new(db.ParentList)
		parent.ClassId = p.ClassId
		parent.Name = p.Name
		_, err := engine.Insert(parent)
		if err != nil {
			log.Log.Error("insert into parent_list :", err)
			return
		}

		childs := p.List
		for _, c := range childs {
			child := new(db.ChildList)
			child.Name = c.Name
			child.ClassId = c.ClassId
			child.ParentId = c.ParentId
			_, err := engine.Insert(child)
			if err != nil {
				log.Log.Error("insert into child_list :", err)
				return
			}
		}
	}

}

func inFood(engine *xorm.Engine, acc *db.AccData, data []byte) {
	foods := acc.Result.ReS.List
	for _, f := range foods {
		food := new(db.FoodList)
		food.Name = f.Name
		food.ClassId = f.ClassId
		food.Pic = f.Pic
		food.Status = f.Status
		food.Tag = f.Tag
		food.Content = f.Content
		food.CookingTime = f.CookingTime
		food.PeopleNum = f.PeopleNum
		food.PrepareTime = f.PrepareTime
		food.SId = f.SId

		_, err := engine.Insert(food)
		if err != nil {
			log.Log.Error("insert into food_list :", err)
			return
		}

		/*var dat map[string]interface{}
		err1 := json.Unmarshal(data, &dat)
		if err1 != nil {
			log.Log.Error("unmarshal dat :", err1)
			continue
		}
		process := dat["result"].(map[string]interface{})["result"].(map[string]interface{})["list"].([]interface{})[0].(map[string]interface{})["process"].(string)
*/
		pros := f.Process
		c := 1
		for _, p := range pros {
			pro := new(db.ProcessList)
			pro.SId = f.SId
			pro.Pic = p.Pic
			pro.PId = c
			pro.PContent = p.PContent
			_, err := engine.Insert(pro)
			if err != nil {
				log.Log.Error("insert into process_list :", err)
				return
			}
			c++
		}

		maters := f.Material
		for _, m := range maters {
			ma := new(db.MaterialList)
			ma.SId = f.SId
			ma.Type = m.Type
			ma.Amount = m.Amount
			ma.MName = m.MName
			_, err := engine.Insert(ma)
			if err != nil {
				log.Log.Error("insert into material_list :", err)
				return
			}
		}
	}
}

func typesSearch(engine *xorm.Engine) {
	ids := make([]db.ChildList, 0)
	err := engine.Cols("class_id").Find(&ids)
	if err != nil {
		log.Log.Error("find class_id :", err)
		return
	}
	count := 0
	for _, id := range ids {
		start := 0
		for {
			url := "https://way.jd.com/jisuapi/byclass?classid=" + id.ClassId + "&start=" + strconv.Itoa(start) + "&num=" + strconv.Itoa(num) + "&appkey=" + getApiKey(engine)
			log.Log.Info("url is :", url)
			data := getRespBody(url)
			acc := new(db.AccData)
			err := json.Unmarshal(data, acc)
			if err != nil {
				log.Log.Error("unmarshal json :", err)
			} else {
				inFood(engine, acc, data)
			}

			start += num
			count++
			page := acc.Result.ReS.Num
			n, _ := strconv.Atoi(page)
			if n < num {
				break
			}
		}
		if count > 640 {
			time.Sleep(time.Hour * 24)
			count = 0
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

	Unmal(engine)
}
