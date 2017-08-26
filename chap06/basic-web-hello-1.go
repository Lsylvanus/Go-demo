package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"goTest/db"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	"ERP_1/log"
)

var e *xorm.Engine
var s *xorm.Session

type Users struct {
	Id            int64     `xorm:"autoincr int(11) pk notnull"`
	Src           string    `xorm:"varchar(5) notnull default ''"`
	UserCode      int64     `xorm:"-"`
	UserNick      string    `xorm:"varchar(20) default ''"`
	Email         string    `xorm:"varchar(30) notnull default ''"`
	Token         string    `xorm:"text null"`
	Mobile        string    `xorm:"varchar(11) default ''"`
	Expire        time.Time `xorm:"datetime notnull default '3000-12-31 23:59:59'"`
	RefreshExpire time.Time `xorm:"datetime notnull default '3000-12-31 23:59:59'"`
	Info          string    `xorm:"text null"`
	Created       time.Time `xorm:"timestamp notnull created"`
	Updated       time.Time `xorm:"timestamp notnull updated"`
}

type Refresh struct {
	Message string       `json:"message"`
	Code    int          `json:"code"`
	Data    RefreshToken `json:"data"`
}

//刷新token
type RefreshToken struct {
	ExpiryTime            int64  `json:"expiry_time"`
	TokenType             string `json:"token_type"`
	Token                 string `json:"access_token"`
	ExpiresIn             int64  `json:"expires_in"`
	Reason                string `json:"reason"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	ClientId              string `json:"client_id"`
	TokenId               string `json:"token_id"`
	MerchantId            string `json:"merchant_id"`
	MerchantUserId        string `json:"merchant_user_id"`
	ClientSecret          string `json:"client_secret"`
	GrantType             string `json:"grant_type"`
}

type EbayToken struct {
	Token                 string `json:"access_token"`
	ExpiresIn             int64  `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	GrantType             string `json:"grant_type"`
}

type Data struct {
	Message  string
	Code     int64
	Method   string
	Previous string
	Next     string
}

//字节转string
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func conn() *time.Location {
	//创建Orm引擎
	var err error
	e, err = xorm.NewEngine("mysql", "root:chenghai3c@/express_delivery?charset=utf8&collation=utf8_general_ci")
	if err != nil {
		log.Log.Info(err.Error())
		return nil
	}
	//测试连接
	errPing := e.Ping()
	if errPing != nil {
		log.Log.Info(errPing.Error())
		return nil
	}
	log.Log.Info("conn success!")

	//事务处理
	s = e.NewSession()
	defer s.Close()
	// add Begin() before any action
	errBegin := s.Begin()
	if errBegin != nil {
		log.Log.Info(errBegin.Error())
		return nil
	}

	//设置日志
	e.ShowSQL(true)
	e.Logger().SetLevel(core.LOG_DEBUG)
	//日志信息保存
	f, err := os.Create("sql_test.log")
	if err != nil {
		log.Log.Info(err.Error())
		return nil
	}
	e.SetLogger(xorm.NewSimpleLogger(f))

	//时区
	cst, _ := time.LoadLocation("Asia/Shanghai")
	return cst
}

func create(beanOrTableName interface{}) bool {
	has, err := e.IsTableExist(beanOrTableName)
	if err != nil {
		log.Log.Info(err.Error())
		return false
	}
	if has {
		log.Log.Info("table is exist.")
		return true
	} else {
		err := e.CreateTables(beanOrTableName)
		if err != nil {
			log.Log.Info("create table failed.")
			return false
		}
		log.Log.Info("create table success.")
	}
	return true
}

func message(code int64, method string, message string, previous string, next string) *Data {
	data := new(Data)
	data.Code = code
	data.Method = method
	data.Message = message
	data.Previous = previous
	data.Next = next
	return data
}

func tmplParse(resp http.ResponseWriter, req *http.Request, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parseFile err :", err.Error())
	}
	err1 := tmpl.Execute(resp, data)
	if err1 != nil {
		fmt.Println("execute data :", err1.Error())
	}
}

func Index(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<b>Hello, Welcome to go web programming...</b><br>"))
}

func AuthInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./auth/form.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err.Error())
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1.Error())
	}
}

func PrintInfo(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Log.Info("parse form :", err.Error())
	}

	src := req.PostFormValue("src")
	log.Log.Info("src : ", src)
	user_code := req.PostFormValue("user_code")
	log.Log.Info("user_code : ", user_code)
	user_nick := req.PostFormValue("user_nick")
	log.Log.Info("user_nick : ", user_nick)
	mobile := req.PostFormValue("mobile")
	log.Log.Info("mobile : ", mobile)
	info := req.PostFormValue("info")
	log.Log.Info("info : ", info)
	secret := req.PostFormValue("secret")
	log.Log.Info("secret : ", secret)
}

func Save(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Log.Info("parse form :", err.Error())
	}

	//获取form表单的值
	src := req.PostFormValue("src")
	log.Log.Info("src : ", src)
	user_code := req.PostFormValue("user_code")
	log.Log.Info("user_code : ", user_code)
	user_nick := req.PostFormValue("user_nick")
	log.Log.Info("user_nick : ", user_nick)
	mobile := req.PostFormValue("mobile")
	log.Log.Info("mobile : ", mobile)
	email := req.PostFormValue("email")
	log.Log.Info("email : ", email)
	info := req.PostFormValue("info")
	log.Log.Info("info : ", info)
	secret := req.PostFormValue("secret")
	log.Log.Info("secret : ", secret)

	//连接数据库
	cst := conn()

	//是否需要创建表
	u := new(Users)
	if b := create(u); !b {
		return
	}

	//处理json，读取token和到期时间
	ref := new(Refresh)
	ebay := new(EbayToken)
	if info != "" {
		//ebay
		if src == "ebay" {
			infoByte := []byte(info)
			err := json.Unmarshal(infoByte, ebay)
			if err != nil {
				log.Log.Info("Unmarshal ref :", err.Error())
				return
			}
		} else {
			//others
			infoByte := []byte(info)
			err := json.Unmarshal(infoByte, ref)
			if err != nil {
				log.Log.Info("Unmarshal ref :", err.Error())
				return
			}
		}
	}

	//user_nick不能重复
	user := new(Users)
	has, err := e.Where("user_nick =?", user_nick).Get(user)
	if err != nil {
		log.Log.Info("select erp_users where user_nick :", err.Error())
		return
	}
	if has {
		data := message(201, "save()", "已有该昵称，请换一个！", "/auth", "")
		tmplParse(resp, req, "./auth/failed.tpl", data)
		d, err := json.Marshal(data)
		if err != nil {
			log.Log.Info("Marshal data :", err.Error())
			return
		}
		log.Log.Info(string(d))
		return
	} else {
		//忽略掉user_code
		if user_code != "" {
			code, _ := strconv.ParseInt(user_code, 10, 0)
			user.UserCode = code
		} else {
			//TODO:
		}

		//email和src都不能为空
		if email == "" || src == "" {
			data := message(101, "save()", "src和email不能为空！", "/auth", "")
			tmplParse(resp, req, "./auth/failed.tpl", data)
			d, err := json.Marshal(data)
			if err != nil {
				log.Log.Info("Marshal ref :", err.Error())
				return
			}
			log.Log.Info(string(d))
			return
		} else {
			user.Src = src
			user.Email = email
		}

		user.Mobile = mobile

		//目前只有wish用户才会用到secret
		if info != "" && secret != "" {
			//wish用户，把secret放在info里面
			ref.Data.ClientSecret = secret
			ref.Data.GrantType = "refresh_token"
			result, err := json.Marshal(ref)
			if err != nil {
				log.Log.Info(err.Error())
				return
			}
			user.Info = byteString(result)
			user.Token = ref.Data.Token
			//时间戳转时间
			expire := time.Unix(ref.Data.ExpiryTime, 0).In(cst)
			user.Expire = expire
			user.RefreshExpire = expire
		} else if secret == "" && info != "" {
			//非wish用户，记录信息
			if user.Src == "ebay" {
				//ebay用户
				user.Info = info
				user.Token = ebay.Token
				expire, _ := time.ParseInLocation("2006-01-02 15:04:05", "3000-12-31 23:59:59", cst)
				user.Expire = expire
				user.RefreshExpire = expire
			} else {
				//others
				user.Token = ref.Data.Token
				expire := time.Unix(ref.Data.ExpiryTime, 0).In(cst)
				user.Expire = expire
				user.RefreshExpire = expire
			}
		} else if secret != "" && info == "" {
			//非wish用户
			data := message(301, "save()", "目前只有wish用户用到secret，如不是wish用户，请在info内填写access_token的返回json！", "/auth", "")
			tmplParse(resp, req, "./auth/failed.tpl", data)
			d, err := json.Marshal(data)
			if err != nil {
				log.Log.Info("Marshal ref :", err.Error())
				return
			}
			log.Log.Info(string(d))
			return
		} else {
			//不填写情况，无token
			//到期时间默认值
			expire, _ := time.ParseInLocation("2006-01-02 15:04:05", "3000-12-31 23:59:59", cst)
			user.Expire = expire
			user.RefreshExpire = expire
		}

		//插入users表
		user.UserNick = user_nick
		aff, err := e.Insert(user) //(&user) panic: reflect: call of reflect.Value.FieldByName on ptr Value
		if err != nil {
			log.Log.Info("insert into erp_users :", err.Error())
			return
		}
		affInt := strconv.FormatInt(aff, 10)
		log.Log.Info("insert into t_users :", affInt)

		data := message(100, "save()", "成功录入授权信息！", "/auth", "")
		tmplParse(resp, req, "./auth/success.tpl", data)
		d, err := json.Marshal(data)
		if err != nil {
			log.Log.Info("Marshal err :", err.Error())
			return
		}
		log.Log.Info(string(d))
	}
}

func ExpressInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./express/express.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err.Error())
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1.Error())
	}
}

func SaveExpress(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Log.Info("parse form :", err.Error())
	}

	//获取form表单的值
	com_type := req.PostFormValue("com_type")
	log.Log.Info("快递公司 : ", com_type) //shunfeng
	post_id := req.PostFormValue("post_id")
	log.Log.Info("快递单号 : ", post_id) //687102589887

	url := "http://www.kuaidi100.com/query?type=" + com_type + "&postid=" + post_id
	r, err := http.Get(url)
	if err != nil {
		log.Log.Info("http client do request :", err.Error())
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Log.Info("io read body :", err.Error())
		return
	}

	//连接数据库
	cst := conn()
	log.Log.Info("now is :", time.Now().In(cst))

	delivery := new(db.Delivery)
	express := new(db.Express)

	err1 := json.Unmarshal(body, delivery)
	if err1 != nil {
		log.Log.Info("json unmarshal :", err1.Error())
		return
	}

	if delivery.Nu == "" {
		delivery.Nu = post_id
		delivery.Com = com_type
		_, err2 := e.Insert(delivery)
		if err2 != nil {
			log.Log.Info("insert into delivery :", err2.Error())
			data := message(201, "SearchExpress()", "记录信息失败！", "/express", "")
			tmplParse(resp, req, "./express/failed.tpl", data)
			return
		}
	} else {
		//nu不能重复
		deli := new(db.Delivery)
		has, err := e.Where("nu =?", delivery.Nu).Get(deli)
		if err != nil {
			log.Log.Info("select delivery where nu :", err.Error())
			return
		}

		if has {
			// todo
		} else {
			_, err2 := e.Insert(delivery)
			if err2 != nil {
				log.Log.Info("insert into delivery :", err2.Error())
				data := message(201, "SearchExpress()", "记录信息失败！", "/express", "")
				tmplParse(resp, req, "./express/failed.tpl", data)
				return
			}
		}
	}

	//nu不能重复
	expr := new(db.Express)
	has, err := e.Where("nu =?", delivery.Nu).Get(expr)
	if err != nil {
		log.Log.Info("select delivery where nu :", err.Error())
		return
	}
	for _, value := range delivery.Data {

		if has {
			//todo
		} else {
			express = &value
			express.Nu = delivery.Nu
			express.FTime, _ = time.ParseInLocation("2006-01-02 15:04:05", value.FTimeStr, cst)
			express.Time, _ = time.ParseInLocation("2006-01-02 15:04:05", value.TimeStr, cst)

			_, err1 := e.Insert(express)
			if err1 != nil {
				log.Log.Info("insert into express :", err1.Error())
				data := message(202, "SearchExpress()", "记录信息失败！", "/express", "")
				tmplParse(resp, req, "./express/failed.tpl", data)
				return
			}
		}
	}

	data := message(100, "SearchExpress()", "记录信息成功！", "/express", "")
	tmplParse(resp, req, "./express/success.tpl", data)
	return
}

func LocInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./loc/loc.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err.Error())
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1.Error())
	}
}

func SaveLoc(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Log.Info("parse form :", err.Error())
	}

	//获取form表单的值
	city := req.PostFormValue("city")
	log.Log.Info("市 : ", city)

	url := "http://gc.ditu.aliyun.com/geocoding?a=" + city
	log.Log.Info("get city url :", url)
	body := getBody(url)

	gc := new(db.GC)
	err1 := json.Unmarshal(body, gc)
	log.Log.Info("gc is :", gc)
	if err1 != nil {
		log.Log.Info("unmarshal gc :", err1.Error())
		return
	}

	urlLoc := "http://ditu.amap.com/service/regeo?longitude=" + strconv.FormatFloat(gc.Lon, 'f', -1, 64) + "&latitude=" + strconv.FormatFloat(gc.Lat, 'f', -1, 64)
	log.Log.Info("get near building url :", urlLoc)
	data := getBody(urlLoc)

	nb := new(db.NearBuilding)
	err2 := json.Unmarshal(data, nb)
	if err2 != nil {
		log.Log.Error("unmarshal nb :", err2.Error())
		return
	}

	//连接数据库
	cst := conn()
	log.Log.Info("now is :", time.Now().In(cst))

	loc := &nb.Data
	_, err8 := e.Insert(loc)
	if err8 != nil {
		log.Log.Error("insert into loc :", err8.Error())
		d := message(201, "SaveLoc()", "记录信息失败！", "/loc", "")
		tmplParse(resp, req, "./loc/failed.tpl", d)
		return
	}

	seaArea := new(db.SeaArea)
	seaArea.LocId = loc.Id
	seaArea.Name = nb.Data.SeaArea.Name
	seaArea.AdCode = nb.Data.SeaArea.AdCode
	_, err3 := e.Insert(seaArea)
	if err3 != nil {
		log.Log.Error("insert into sea_area :", err3.Error())
		d := message(202, "SaveLoc()", "记录信息失败！", "/loc", "")
		tmplParse(resp, req, "./loc/failed.tpl", d)
		return
	}

	for _, v := range nb.Data.CrossList {
		crossList := new(db.CrossList)
		crossList.LocId = loc.Id
		crossList.Name = v.Name
		crossList.CrossId = v.CrossId
		crossList.Direction = v.Direction
		crossList.Distance = v.Distance
		crossList.Latitude = v.Latitude
		crossList.Level = v.Level
		crossList.Longitude = v.Longitude
		crossList.Weight = v.Weight
		crossList.Width = v.Width

		_, err4 := e.Insert(crossList)
		if err4 != nil {
			log.Log.Error("insert into cross_list :", err4.Error())
			d := message(203, "SaveLoc()", "记录信息失败！", "/loc", "")
			tmplParse(resp, req, "./loc/failed.tpl", d)
			return
		}
	}

	for _, v := range nb.Data.RoadList {
		roadList := new(db.RoadList)
		roadList.LocId = loc.Id
		roadList.Width = v.Width
		roadList.Longitude = v.Longitude
		roadList.Level = v.Level
		roadList.Latitude = v.Latitude
		roadList.Distance = v.Distance
		roadList.Direction = v.Direction
		roadList.RoadId = v.RoadId
		roadList.Name = v.Name

		_, err5 := e.Insert(roadList)
		if err5 != nil {
			log.Log.Error("insert into road_list :", err5.Error())
			d := message(204, "SaveLoc()", "记录信息失败！", "/loc", "")
			tmplParse(resp, req, "./loc/failed.tpl", d)
			return
		}
	}

	for _, v := range nb.Data.PoiList {
		poiList := new(db.PoiList)
		poiList.LocId = loc.Id
		poiList.Name = v.Name
		poiList.Direction = v.Direction
		poiList.Distance = v.Distance
		poiList.Latitude = v.Latitude
		poiList.Longitude = v.Longitude
		poiList.Weight = v.Weight
		poiList.Type = v.Type
		poiList.TypeCode = v.TypeCode
		poiList.Address = v.Address
		poiList.PoiId = v.PoiId
		poiList.Tel = v.Tel

		_, err6 := e.Insert(poiList)
		if err6 != nil {
			log.Log.Error("insert into poi_list :", err6.Error())
			d := message(205, "SaveLoc()", "记录信息失败！", "/loc", "")
			tmplParse(resp, req, "./loc/failed.tpl", d)
			return
		}
		for _, vs := range v.Entrances {
			entrances := new(db.Entrances)
			entrances.LocId = loc.Id
			entrances.PoiListId = poiList.Id
			entrances.Longitude = vs.Longitude
			entrances.Latitude = vs.Latitude

			_, err7 := e.Insert(entrances)
			if err7 != nil {
				log.Log.Error("insert into entrances :", err7.Error())
				d := message(206, "SaveLoc()", "记录信息失败！", "/loc", "")
				tmplParse(resp, req, "./loc/failed.tpl", d)
				return
			}
		}
	}

	d := message(100, "SaveLoc()", "记录信息成功！", "/loc", "")
	tmplParse(resp, req, "./loc/success.tpl", d)
	return

}

func getBody(url string) []byte {
	r, err := http.Get(url)
	if err != nil {
		log.Log.Info("http client do request :", err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Log.Info("io read body :", err.Error())
		return nil
	}
	return body
}

func main() {
	log.LogInit()
	http.HandleFunc("/", Index)
	http.HandleFunc("/auth", AuthInfo)
	http.HandleFunc("/info/print", PrintInfo)
	http.HandleFunc("/auth/user/save", Save)
	http.HandleFunc("/express", ExpressInfo)
	http.HandleFunc("/express/save", SaveExpress)
	http.HandleFunc("/loc", LocInfo)
	http.HandleFunc("/loc/save", SaveLoc)
	http.ListenAndServe("192.168.10.197:8888", nil)
}
