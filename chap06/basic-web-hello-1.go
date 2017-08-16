package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
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
	e, err = xorm.NewEngine("mysql", "root:chenghai3c@/wish_catch?charset=utf8&collation=utf8_general_ci")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	//测试连接
	errPing := e.Ping()
	if errPing != nil {
		log.Println(errPing.Error())
		return nil
	}
	log.Println("conn success!")

	//事务处理
	s = e.NewSession()
	defer s.Close()
	// add Begin() before any action
	errBegin := s.Begin()
	if errBegin != nil {
		log.Println(errBegin.Error())
		return nil
	}

	//设置日志
	e.ShowSQL(true)
	e.Logger().SetLevel(core.LOG_DEBUG)
	//日志信息保存
	f, err := os.Create("sql_test.log")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	e.SetLogger(xorm.NewSimpleLogger(f))

	// 映射规则
	tMapper := core.NewPrefixMapper(core.SnakeMapper{}, "erp_")
	e.SetTableMapper(tMapper)

	//时区
	cst, _ := time.LoadLocation("Asia/Shanghai")
	return cst
}

func create(beanOrTableName interface{}) bool {
	has, err := e.IsTableExist(beanOrTableName)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if has {
		log.Println("table is exist.")
		return true
	} else {
		err := e.CreateTables(beanOrTableName)
		if err != nil {
			log.Println("create table failed.")
			return false
		}
		log.Println("create table success.")
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
	tmpl, err := template.ParseFiles("./form.tpl")
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
		log.Println("parse form :", err.Error())
	}

	src := req.PostFormValue("src")
	log.Println("src : ", src)
	user_code := req.PostFormValue("user_code")
	log.Println("user_code : ", user_code)
	user_nick := req.PostFormValue("user_nick")
	log.Println("user_nick : ", user_nick)
	mobile := req.PostFormValue("mobile")
	log.Println("mobile : ", mobile)
	info := req.PostFormValue("info")
	log.Println("info : ", info)
	secret := req.PostFormValue("secret")
	log.Println("secret : ", secret)
}

func Save(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println("parse form :", err.Error())
	}

	//获取form表单的值
	src := req.PostFormValue("src")
	log.Println("src : ", src)
	user_code := req.PostFormValue("user_code")
	log.Println("user_code : ", user_code)
	user_nick := req.PostFormValue("user_nick")
	log.Println("user_nick : ", user_nick)
	mobile := req.PostFormValue("mobile")
	log.Println("mobile : ", mobile)
	email := req.PostFormValue("email")
	log.Println("email : ", email)
	info := req.PostFormValue("info")
	log.Println("info : ", info)
	secret := req.PostFormValue("secret")
	log.Println("secret : ", secret)

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
				log.Println("Unmarshal ref :", err.Error())
				return
			}
		} else {
			//others
			infoByte := []byte(info)
			err := json.Unmarshal(infoByte, ref)
			if err != nil {
				log.Println("Unmarshal ref :", err.Error())
				return
			}
		}
	}

	//user_nick不能重复
	user := new(Users)
	has, err := e.Where("user_nick =?", user_nick).Get(user)
	if err != nil {
		log.Println("select erp_users where user_nick :", err.Error())
		return
	}
	if has {
		data := message(201, "save()", "已有该昵称，请换一个！", "/auth", "")
		tmplParse(resp, req, "./failed.tpl", data)
		d, err := json.Marshal(data)
		if err != nil {
			log.Println("Marshal data :", err.Error())
			return
		}
		log.Println(string(d))
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
			tmplParse(resp, req, "./failed.tpl", data)
			d, err := json.Marshal(data)
			if err != nil {
				log.Println("Marshal ref :", err.Error())
				return
			}
			log.Println(string(d))
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
				log.Println(err.Error())
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
			tmplParse(resp, req, "./failed.tpl", data)
			d, err := json.Marshal(data)
			if err != nil {
				log.Println("Marshal ref :", err.Error())
				return
			}
			log.Println(string(d))
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
			log.Println("insert into erp_users :", err.Error())
			return
		}
		affInt := strconv.FormatInt(aff, 10)
		log.Println("insert into t_users :", affInt)

		data := message(100, "save()", "成功录入授权信息！", "/auth", "")
		tmplParse(resp, req, "./success.tpl", data)
		d, err := json.Marshal(data)
		if err != nil {
			log.Println("Marshal err :", err.Error())
			return
		}
		log.Println(string(d))
	}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/auth", AuthInfo)
	http.HandleFunc("/info/print", PrintInfo)
	http.HandleFunc("/auth/user/save", Save)
	http.ListenAndServe("192.168.10.197:8888", nil)
}
