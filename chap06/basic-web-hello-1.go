package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/labstack/gommon/log"
	"goTest/db"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	Lat      float64
	Lon      float64
	Previous string
	Next     string
}

type HtmlAttr struct {
	Author      string
	Description string
	Title       string
	HeadHref    string
	PostAction  string
	LabelText   string
	Placeholder string
	TextId      string
	TextName    string
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
		fmt.Println(err)
		return nil
	}
	//测试连接
	errPing := e.Ping()
	if errPing != nil {
		fmt.Println(errPing)
		return nil
	}
	fmt.Println("conn success!")

	//事务处理
	s = e.NewSession()
	defer s.Close()
	// add Begin() before any action
	errBegin := s.Begin()
	if errBegin != nil {
		fmt.Println(errBegin)
		return nil
	}

	//设置日志
	e.ShowSQL(true)
	e.Logger().SetLevel(core.LOG_DEBUG)
	//日志信息保存
	f, err := os.Create("sql_test.log")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	e.SetLogger(xorm.NewSimpleLogger(f))

	//时区
	cst, _ := time.LoadLocation("Asia/Shanghai")
	e.TZLocation = cst
	return cst
}

func create(beanOrTableName interface{}) bool {
	has, err := e.IsTableExist(beanOrTableName)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if has {
		fmt.Println("table is exist.")
		return true
	} else {
		err := e.CreateTables(beanOrTableName)
		if err != nil {
			fmt.Println("create table failed.")
			return false
		}
		fmt.Println("create table success.")
	}
	return true
}

func message(code int64, method string, message string, lon float64, lat float64, previous string, next string) *Data {
	data := new(Data)
	data.Code = code
	data.Method = method
	data.Message = message
	data.Lon = lon
	data.Lat = lat
	data.Previous = previous
	data.Next = next
	return data
}

func tmplParse(resp http.ResponseWriter, req *http.Request, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parseFile err :", err)
	}
	err1 := tmpl.Execute(resp, data)
	if err1 != nil {
		fmt.Println("execute data :", err1)
	}
}

func Index(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<b>Hello, Welcome to go web programming...</b><br>"))
}

func AuthInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./auth/form.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err)
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1)
	}
}

func PrintInfo(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("parse form :", err)
	}

	src := req.PostFormValue("src")
	fmt.Println("src : ", src)
	user_code := req.PostFormValue("user_code")
	fmt.Println("user_code : ", user_code)
	user_nick := req.PostFormValue("user_nick")
	fmt.Println("user_nick : ", user_nick)
	mobile := req.PostFormValue("mobile")
	fmt.Println("mobile : ", mobile)
	info := req.PostFormValue("info")
	fmt.Println("info : ", info)
	secret := req.PostFormValue("secret")
	fmt.Println("secret : ", secret)
}

func Save(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("parse form :", err)
	}

	//获取form表单的值
	src := req.PostFormValue("src")
	fmt.Println("src : ", src)
	user_code := req.PostFormValue("user_code")
	fmt.Println("user_code : ", user_code)
	user_nick := req.PostFormValue("user_nick")
	fmt.Println("user_nick : ", user_nick)
	mobile := req.PostFormValue("mobile")
	fmt.Println("mobile : ", mobile)
	email := req.PostFormValue("email")
	fmt.Println("email : ", email)
	info := req.PostFormValue("info")
	fmt.Println("info : ", info)
	secret := req.PostFormValue("secret")
	fmt.Println("secret : ", secret)

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
				fmt.Println("Unmarshal ref :", err)
				return
			}
		} else {
			//others
			infoByte := []byte(info)
			err := json.Unmarshal(infoByte, ref)
			if err != nil {
				fmt.Println("Unmarshal ref :", err)
				return
			}
		}
	}

	//user_nick不能重复
	user := new(Users)
	has, err := e.Where("user_nick =?", user_nick).Get(user)
	if err != nil {
		fmt.Println("select erp_users where user_nick :", err)
		return
	}
	if has {
		data := message(201, "save()", "已有该昵称，请换一个！", 0, 0, "/auth", "")
		tmplParse(resp, req, "./auth/failed.tpl", data)
		d, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Marshal data :", err)
			return
		}
		fmt.Println(string(d))
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
			data := message(101, "save()", "src和email不能为空！", 0, 0, "/auth", "")
			tmplParse(resp, req, "./auth/failed.tpl", data)
			d, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Marshal ref :", err)
				return
			}
			fmt.Println(string(d))
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
				fmt.Println(err)
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
			data := message(301, "save()", "目前只有wish用户用到secret，如不是wish用户，请在info内填写access_token的返回json！", 0, 0, "/auth", "")
			tmplParse(resp, req, "./auth/failed.tpl", data)
			d, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Marshal ref :", err)
				return
			}
			fmt.Println(string(d))
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
			fmt.Println("insert into erp_users :", err)
			return
		}
		affInt := strconv.FormatInt(aff, 10)
		fmt.Println("insert into t_users :", affInt)

		data := message(100, "save()", "成功录入授权信息！", 0, 0, "/auth", "")
		tmplParse(resp, req, "./auth/success.tpl", data)
		d, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Marshal err :", err)
			return
		}
		fmt.Println(string(d))
	}
}

func ExpressInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./express/express.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err)
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1)
	}
}

func SaveExpress(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("parse form :", err)
	}

	//获取form表单的值
	com_type := req.PostFormValue("com_type")
	fmt.Println("快递公司 : ", com_type) //shunfeng	ems	yunda
	post_id := req.PostFormValue("post_id")
	fmt.Println("快递单号 : ", post_id) //687102589887	1182457917532	3840700256014

	url := "http://www.kuaidi100.com/query?type=" + strings.TrimSpace(com_type) + "&postid=" + strings.TrimSpace(post_id)
	fmt.Println("express url is :", url)
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("http client do request :", err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("io read body :", err)
		return
	}

	//连接数据库
	cst := conn()
	fmt.Println("now is :", time.Now().In(cst))

	delivery := new(db.Delivery)
	express := new(db.Express)

	err1 := json.Unmarshal(body, delivery)
	if err1 != nil {
		fmt.Println("json unmarshal :", err1)
		return
	}

	if delivery.Nu == "" {
		delivery.Nu = post_id
		delivery.Com = com_type
		_, err2 := e.Insert(delivery)
		if err2 != nil {
			fmt.Println("insert into delivery :", err2)
			data := message(201, "insert into delivery with no data.", err2.Error(), 0, 0, "/express", "")
			tmplParse(resp, req, "./express/failed.tpl", data)
			return
		}
	} else {
		//nu不能重复
		deli := new(db.Delivery)
		has, err := e.Where("nu =?", delivery.Nu).Get(deli)
		if err != nil {
			fmt.Println("select delivery where nu :", err)
			return
		}

		if has {
			// todo
		} else {
			_, err2 := e.Insert(delivery)
			if err2 != nil {
				fmt.Println("insert into delivery :", err2)
				data := message(201, "insert into delivery.", err2.Error(), 0, 0, "/express", "")
				tmplParse(resp, req, "./express/failed.tpl", data)
				return
			}
		}
	}

	//nu不能重复
	expr := new(db.Express)
	has, err := e.Where("nu =?", delivery.Nu).Get(expr)
	if err != nil {
		fmt.Println("select delivery where nu :", err)
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
				fmt.Println("insert into express :", err1)
				data := message(202, "insert into express.", err1.Error(), 0, 0, "/express", "")
				tmplParse(resp, req, "./express/failed.tpl", data)
				return
			}
		}
	}

	data := message(100, "Successful!", "记录信息成功！", 0, 0, "/express", "")
	tmplParse(resp, req, "./express/success.tpl", data)

	return
}

func LocInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./loc/loc.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err)
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1)
	}
}

func ShowLoc(resp http.ResponseWriter, req *http.Request) {

}

func SaveLoc(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("parse form :", err)
	}

	//获取form表单的值
	city := req.PostFormValue("city")
	fmt.Println("市 : ", city)

	url := "http://gc.ditu.aliyun.com/geocoding?a=" + city
	fmt.Println("get city url :", url)
	body := getBody(url)

	gc := new(db.GC)
	err1 := json.Unmarshal(body, gc)
	if err1 != nil {
		fmt.Println("unmarshal gc :", err1)
		return
	}

	urlLoc := "http://ditu.amap.com/service/regeo?longitude=" + strconv.FormatFloat(gc.Lon, 'f', -1, 64) + "&latitude=" + strconv.FormatFloat(gc.Lat, 'f', -1, 64)
	fmt.Println("get near building url :", urlLoc)
	data := getBody(urlLoc)

	nb := new(db.NearBuilding)
	err2 := json.Unmarshal(data, nb)
	if err2 != nil {
		fmt.Println("unmarshal nb :", err2)
		return
	}

	//连接数据库
	cst := conn()
	fmt.Println("now is :", time.Now().In(cst))

	loc := &nb.Data
	loc.Gc = city
	_, err8 := e.Insert(loc)
	if err8 != nil {
		fmt.Println("insert into loc :", err)
		d := message(201, "insert into loc.", err.Error(), 0, 0, "/loc", "")
		tmplParse(resp, req, "./loc/failed.tpl", d)
		return
	}

	seaArea := new(db.SeaArea)
	seaArea.LocId = loc.Id
	seaArea.Name = nb.Data.SeaArea.Name
	seaArea.AdCode = nb.Data.SeaArea.AdCode
	_, err3 := e.Insert(seaArea)
	if err3 != nil {
		fmt.Println("insert into sea_area :", err)
		d := message(202, "insert into sea_list.", err.Error(), 0, 0, "/loc", "")
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
			fmt.Println("insert into cross_list :", err4)
			d := message(203, "insert into cross_list.", err4.Error(), 0, 0, "/loc", "")
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
			fmt.Println("insert into road_list :", err5)
			d := message(204, "insert into road_list.", err5.Error(), 0, 0, "/loc", "")
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
			fmt.Println("insert into poi_list :", err6)
			d := message(205, "insert into poi_list.", err6.Error(), 0, 0, "/loc", "")
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
				fmt.Println("insert into entrances :", err7)
				d := message(206, "insert into entrances.", err7.Error(), 0, 0, "/loc", "")
				tmplParse(resp, req, "./loc/failed.tpl", d)
				return
			}
		}
	}

	d := message(100, "Successful!", "记录信息成功！", gc.Lon, gc.Lat, "/loc", "")
	tmplParse(resp, req, "./loc/success.tpl", d)
	return
}

func ZoneMusicInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./qzone/qzone.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err)
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1)
	}
}

func SaveQZoneMusic(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("parse form :", err)
	}

	//获取form表单的值
	qq := req.PostFormValue("qq")
	fmt.Println("QQ : ", qq)

	url := "http://qzone-music.qq.com/fcg-bin/cgi_playlist_xml.fcg?uin=" + qq + "&json=1&g_tk=1916754934"
	fmt.Println("get qzone music url :", url)
	body := getBody(url)

	data, _ := ParseUnformattedJson(body)
	fmt.Println("data is :", data)
}

func WeatherInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./weather/weather.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err)
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1)
	}
}

func SaveWeatherInfo(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("parse form :", err)
	}

	//获取form表单的值
	city := req.PostFormValue("city")
	fmt.Println("city : ", city)

	url := "http://www.sojson.com/open/api/weather/json.shtml?city=" + strings.TrimSpace(city)
	fmt.Println("get weather url :", url)
	body := getBody(url)

	reback := new(db.ReBack)
	err1 := json.Unmarshal(body, reback)
	if err1 != nil {
		fmt.Println("unmarshal body :", err1)
	}

	//连接数据库
	cst := conn()
	fmt.Println("now is :", time.Now().In(cst))

	weather := &reback.Data
	weather.City = reback.City
	fmt.Println("weather is :", weather)
	_, err3 := e.Insert(weather)
	if err3 != nil {
		fmt.Println("insert into weather :", err3)
		d := message(201, "insert into weather.", err3.Error(), 0, 0, "/weather", "")
		tmplParse(resp, req, "./weather/failed.tpl", d)
		return
	}

	yesterday := new(db.Yesterday)
	yesterday.CityId = weather.Id
	yesterday.Type = reback.Data.Yesterday.Type
	yesterday.Date = reback.Data.Yesterday.Date
	yesterday.Aqi = reback.Data.Yesterday.Aqi
	yesterday.Fl = reback.Data.Yesterday.Fl
	yesterday.Fx = reback.Data.Yesterday.Fx
	yesterday.High = reback.Data.Yesterday.High
	yesterday.Low = reback.Data.Yesterday.Low
	yesterday.Sunrise = reback.Data.Yesterday.Sunrise
	yesterday.Sunset = reback.Data.Yesterday.Sunset
	yesterday.Notice = reback.Data.Yesterday.Notice
	_, err4 := e.Insert(yesterday)
	if err4 != nil {
		fmt.Println("insert into yesterday :", err4)
		d := message(202, "insert into yesterday.", err4.Error(), 0, 0, "/weather", "")
		tmplParse(resp, req, "./weather/failed.tpl", d)
		return
	}

	for _, v := range reback.Data.Forecast {
		forecast := new(db.Forecast)
		forecast.CityId = weather.Id
		forecast.Type = v.Type
		forecast.Date = v.Date
		forecast.Aqi = v.Aqi
		forecast.Fl = v.Fl
		forecast.Fx = v.Fx
		forecast.High = v.High
		forecast.Low = v.Low
		forecast.Sunrise = v.Sunrise
		forecast.Sunset = v.Sunset
		forecast.Notice = v.Notice
		_, err4 := e.Insert(forecast)
		if err4 != nil {
			fmt.Println("insert into forecast :", err4)
			d := message(203, "insert into forecast.", err4.Error(), 0, 0, "/weather", "")
			tmplParse(resp, req, "./weather/failed.tpl", d)
			return
		}
	}
	d := message(100, "Successful!", "记录信息成功！", 0, 0, "/weather", "")
	tmplParse(resp, req, "./weather/success.tpl", d)
	return
}

func MusicInfo(resp http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("./music/music.tpl")
	if err != nil {
		fmt.Println("parseFile err :", err)
	}
	err1 := tmpl.Execute(resp, nil)
	if err1 != nil {
		fmt.Println("execute data :", err1)
	}
}

func SaveMusicInfo(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("parse form :", err)
	}

	//获取form表单的值
	keyword := req.PostFormValue("keyword")
	fmt.Println("keyword : ", keyword)

	url := "http://sug.music.baidu.com/info/suggestion?word=" + keyword + "&version=2&from=0"
	data := getBody(url)

	songsData := new(db.SearchMusic)
	err1 := json.Unmarshal(data, songsData)
	if err1 != nil {
		fmt.Println("json unmarshal songs data :", err1)
		return
	}

	//连接数据库
	cst := conn()
	fmt.Println("now is :", time.Now().In(cst))

	songs := songsData.Data.Songs
	artists := songsData.Data.Artist
	albums := songsData.Data.Album
	for _, song := range songs {
		sg := new(db.Song)
		has, err := e.Where("songid =?", song.Songid).Get(sg)
		if err != nil {
			fmt.Println("select from song :", err)
			return
		}
		if !has {
			s := new(db.Song)
			s.Info = song.Info
			s.Weight = song.Weight
			s.Artistname = song.Artistname
			s.BitrateFee = song.BitrateFee
			s.EncryptedSongid = song.EncryptedSongid
			s.HasMv = song.HasMv
			s.ResourceProvider = song.ResourceProvider
			s.ResourceTypeExt = song.ResourceTypeExt
			s.Songid = song.Songid
			s.Songname = song.Songname
			s.YyrArtist = song.YyrArtist

			_, err := e.Insert(s)
			if err != nil {
				fmt.Println("insert into song :", err)
				d := message(201, "insert into song.", err.Error(), 0, 0, "/music", "")
				tmplParse(resp, req, "./music/failed.tpl", d)
				return
			}
		}
	}

	for _, artist := range artists {
		ar := new(db.Artist)
		has, err := e.Where("artistid =?", artist.Artistid).Get(ar)
		if err != nil {
			fmt.Println("select from artist :", err)
			return
		}
		if !has {
			a := new(db.Artist)
			a.YyrArtist = artist.YyrArtist
			a.Artistname = artist.Artistname
			a.Weight = artist.Weight
			a.Artistid = artist.Artistid
			a.Artistpic = artist.Artistpic
			_, err := e.Insert(a)
			if err != nil {
				fmt.Println("insert into artist :", err)
				d := message(202, "insert into artist.", err.Error(), 0, 0, "/music", "")
				tmplParse(resp, req, "./music/failed.tpl", d)
				return
			}
		}
	}

	for _, album := range albums {
		am := new(db.Album)
		has, err := e.Where("albumid =?", album.Albumid).Get(am)
		if err != nil {
			fmt.Println("select from album :", err)
			return
		}
		if !has {
			al := new(db.Album)
			al.Artistpic = album.Artistpic
			al.Weight = album.Weight
			al.Artistname = album.Artistname
			al.ResourceTypeExt = album.ResourceTypeExt
			al.Albumid = album.Albumid
			al.Albumname = album.Albumname
			_, err := e.Insert(al)
			if err != nil {
				fmt.Println("insert into album :", err)
				d := message(203, "insert into album.", err.Error(), 0, 0, "/music", "")
				tmplParse(resp, req, "./music/failed.tpl", d)
				return
			}
		}
	}
	d := message(200, "insert into table successful.", "记录信息成功！", 0, 0, "/music", "")
	tmplParse(resp, req, "./music/success.tpl", d)
	return
}

// 百度视听返回的json格式不确定，字段类型有时是int，有时是string
// 用interface{}来代替结构体
// 只拿确定类型的数据(interface nil not string)
// todo: 还是有问题，会卡住跑不下去。第一次跑没问题，第二次就卡住，也不抛出错误。重新运行也不行
func SaveDownLink(resp http.ResponseWriter, req *http.Request) {

	/*song := new(db.Song)
	rows, err := e.Count(song)
	if err != nil {
		fmt.Println("select count(1) from song :", err)
		return
	}*/
	songIds := make([]db.Song, 0)
	err2 := e.SQL("SELECT * FROM song where songid not in (SELECT song_id FROM song_list)").Find(&songIds)
	if err2 != nil {
		fmt.Println("select songid from song :", err2)
		return
	}
	/*sql := "SELECT songid FROM song where songid not in (SELECT DISTINCT song_id FROM song_list)"
	songIdsList, err := e.Query(sql)
	if err != nil {
		fmt.Println("select distinct from song and song_list :", err)
		return
	}*/

	for _, song := range songIds {
		fmt.Println("cur id is :", song.Songid)
		sg := new(db.SongList)
		has, err := e.Where("song_id =?", song.Songid).Get(sg)
		if err != nil {
			fmt.Println("select from song_list :", err)
			return
		}
		if !has {
			linkUrl := "http://music.baidu.com/data/music/fmlink?songIds=" + song.Songid + "&type=flac"
			fmt.Println("downlink url is :", linkUrl)
			dataLink := getBody(linkUrl)
			// todo：输出完link就卡住了
			/*downs := new(db.DownLink)
			err := json.Unmarshal(dataLink, downs)
			if err != nil {
				fmt.Println("unmarshal downs :", err)
				return
			}*/
			var dat map[string]interface{}
			err = json.Unmarshal(dataLink, &dat)
			if err != nil {
				fmt.Println("unmarshal interface :", err)
				return
			}

			songList := dat["data"].(map[string]interface{})["songList"].([]interface{})[0]

			list := new(db.SongList)

			album := new(db.Album)
			if songList.(map[string]interface{})["albumName"] != nil {
				b, err := e.Where("albumname =?", songList.(map[string]interface{})["albumName"].(string)).Get(album)
				if err != nil {
					fmt.Println("select from album where albumname :", err)
					return
				}
				if b {
					list.AlbumId = album.Albumid
					list.AlbumName = album.Albumname
				} else {
					list.ArtistId = ""
					list.ArtistName = ""
				}
			}

			list.SongId = song.Songid
			list.SongName = song.Songname


			artist := new(db.Artist)
			c, err := e.Where("artistname =?", songList.(map[string]interface{})["artistName"].(string)).Get(artist)
			if err != nil {
				fmt.Println("select from artist where artistname :", err)
				return
			}
			if c {
				list.ArtistId = songList.(map[string]interface{})["artistId"].(string)
			} else {
				list.ArtistId = ""
			}
			list.ArtistName = songList.(map[string]interface{})["artistName"].(string)

			list.Time = songList.(map[string]interface{})["time"].(float64)
			list.CopyType = songList.(map[string]interface{})["copyType"].(float64)
			list.LinkCode = songList.(map[string]interface{})["linkCode"].(float64)

			list.QueryId = songList.(map[string]interface{})["queryId"].(string)
			list.Version = songList.(map[string]interface{})["version"].(string)
			list.Format = songList.(map[string]interface{})["format"].(string)
			list.RelateStatus = songList.(map[string]interface{})["relateStatus"].(string)
			list.LrcLink = songList.(map[string]interface{})["lrcLink"].(string)
			list.ShowLink = songList.(map[string]interface{})["showLink"].(string)
			list.SongLink = songList.(map[string]interface{})["songLink"].(string)
			list.Source = songList.(map[string]interface{})["source"].(string)
			list.ResourceType = songList.(map[string]interface{})["resourceType"].(string)

			if songList.(map[string]interface{})["songPicBig"] != nil {
				list.SongPicBig = songList.(map[string]interface{})["songPicBig"].(string)
			}
			if songList.(map[string]interface{})["songPicRadio"] != nil {
				list.SongPicRadio = songList.(map[string]interface{})["songPicRadio"].(string)
			}
			if songList.(map[string]interface{})["songPicSmall"] != nil {
				list.SongPicSmall = songList.(map[string]interface{})["songPicSmall"].(string)
			}

			_, err1 := e.Insert(list)
			if err1 != nil {
				fmt.Println("insert into list :", err1)
				d := message(204, "insert into song_list.", err.Error(), 0, 0, "/music", "")
				tmplParse(resp, req, "./music/failed.tpl", d)
				return
			}
			fmt.Println("insert finished...")
		}
	}
	time.Sleep(time.Second * 3)
	d := message(200, "insert into song_list successful.", "记录信息成功！", 0, 0, "/music", "")
	tmplParse(resp, req, "./music/success.tpl", d)
	return
}

func getBody(url string) []byte {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("http client do request :", err)
		return nil
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("io read body :", err)
		return nil
	}
	return body
}

func ParseUnformattedJson(body []byte) (map[string]interface{}, error) {
	var dat map[string]interface{}
	str1 := strings.Replace(string(body), "jsonCallback({", "{", -1)
	str2 := strings.Replace(str1, "})", "}", -1)
	newbody := []byte(str2)

	arr := strings.Split(str2, ",")
	for k, v := range arr {
		fmt.Printf("key is :%v, value is :%v\n", k, v)
	}

	err := json.Unmarshal(newbody, &dat)
	if err != nil {
		fmt.Println("unmarshal qzone :", err)
		return nil, err
	}
	return dat, nil
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/auth", AuthInfo)
	http.HandleFunc("/info/print", PrintInfo)
	http.HandleFunc("/auth/user/save", Save)
	http.HandleFunc("/express", ExpressInfo)
	http.HandleFunc("/express/save", SaveExpress)
	http.HandleFunc("/loc", LocInfo)
	http.HandleFunc("/loc/show", ShowLoc)
	http.HandleFunc("/loc/save", SaveLoc)
	http.HandleFunc("/qzone", ZoneMusicInfo)
	http.HandleFunc("/qzone/music/save", SaveQZoneMusic)
	http.HandleFunc("/weather", WeatherInfo)
	http.HandleFunc("/weather/save", SaveWeatherInfo)
	http.HandleFunc("/music", MusicInfo)
	http.HandleFunc("/music/search", SaveMusicInfo)
	http.HandleFunc("/music/link/save", SaveDownLink)
	log.Fatal(http.ListenAndServe("192.168.10.197:8098", nil))
}
