package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"path/filepath"
	"runtime"
	"path"
	"log"
)

var start string = "0"

const (
	acc_key = "b637d49acf7c9644fc7d39d11e894fee"
	num     = 5
	filename= "out.txt"
)

type Deli struct {
	Message   string `json:"message"`
	Nu        string `json:"nu"`
	IsCheck   string `json:"ischeck" xorm:"'ischeck'"`
	Condition string `json:"condition"`
	Com       string `json:"com"`
	Status    string `json:"status"`
	State     string `json:"state"`
	Data      []Expr `json:"data" xorm:"-"`
}
type Expr struct {
	Nu       string
	TimeStr  string `json:"time" xorm:"-"`
	Time     time.Time
	FTimeStr string    `json:"ftime" xorm:"-"`
	FTime    time.Time `xorm:"'ftime'"`
	Context  string    `json:"context"`
	Location string    `json:"location"`
}

type WeBack struct {
	Code   string   `json:"code"`
	Charge bool     `json:"charge"`
	Msg    string   `json:"msg"`
	Result WeResult `json:"result"`
}
type WeResult struct {
	HeWeather5 []HeWeather5 `json:"HeWeather5"`
}
type HeWeather5 struct {
	Now            Now              `json:"now"`
	Suggestion     Suggestion       `json:"suggestion"`
	Aqi            Aqi              `json:"aqi"`
	Basic          Basic            `json:"basic"`
	DailyForecast  []DailyForecast  `json:"daily_forecast"`
	HourlyForecast []HourlyForecast `json:"hourly_forecast"`
	Status         string           `json:"status"`
}
type Suggestion struct {
	Uv struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
	Cw struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
	Trav struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
	Air struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
	Comf struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
	Drsg struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
	Sport struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
	Flu struct {
		Txt string `json:"txt"`
		Brf string `json:"brf"`
	}
}
type Aqi struct {
	City C `json:"city"`
}
type C struct {
	No2  string `json:"no2"`
	O3   string `json:"o3"`
	Pm25 string `json:"pm25"`
	Qlty string `json:"qlty"`
	So2  string `json:"so2"`
	Aqi  string `json:"aqi"`
	Pm10 string `json:"pm10"`
	Co   string `json:"co"`
}
type Basic struct {
	City   string `json:"city"`
	Update Update `json:"update"`
	Lon    string `json:"lon"`
	Id     string `json:"id"`
	Cnty   string `json:"cnty"`
	Lat    string `json:"lat"`
}
type Update struct {
	Loc string `json:"loc"`
	Utc string `json:"utc"`
}
type DailyForecast struct {
	Date  string `json:"date"`
	Pop   string `json:"pop"`
	Hum   string `json:"hum"`
	Uv    string `json:"uv"`
	Vis   string `json:"vis"`
	Astro Astro  `json:"astro"`
	Pres  string `json:"press"`
	Pcpn  string `json:"pcpn"`
	Tmp   Tmp    `json:"tmp"`
	Cond  WCond  `json:"cond"`
	Wind  Wind   `json:"wind"`
}
type Astro struct {
	Ss string `json:"ss"`
	Mr string `json:"mr"`
	Ms string `json:"ms"`
	Sr string `json:"sr"`
}
type Tmp struct {
	Min string `json:"min"`
	Max string `json:"max"`
}
type WCond struct {
	Txtn  string `json:"txt_n"`
	Coden string `json:"code_n"`
	Coded string `json:"code_d"`
	Txtd  string `json:"txt_d"`
}
type HourlyForecast struct {
	Date string `json:"date"`
	Pop  string `json:"pop"`
	Hum  string `json:"hum"`
	Pres string `json:"press"`
	Tmp  string `json:"tmp"`
	Cond Cond   `json:"cond"`
	Wind Wind   `json:"wind"`
}
type Now struct {
	Hum  string `json:"hum"`
	Vis  string `json:"vis"`
	Pres string `json:"pres"`
	Pcpn string `json:"pcpn"`
	Fl   string `json:"fl"`
	Tmp  string `json:"tmp"`
	Cond Cond   `json:"cond"`
	Wind Wind   `json:"wind"`
}
type Cond struct {
	Txt  string `json:"txt"`
	Code string `json:"code"`
}
type Wind struct {
	Sc  string `json:"sc"`
	Spd string `json:"spd"`
	Deg string `json:"deg"`
	Dir string `json:"dir"`
}

type ReBack struct {
	Message string  `json:"message"`
	Status  int     `json:"status"`
	City    string  `json:"city"`
	Count   int     `json:"count"`
	Data    Weather `json:"data"`
}
type Weather struct {
	City      string
	Shidu     string     `json:"shidu"`
	Pm25      float64    `json:"pm25"`
	Pm10      float64    `json:"pm10"`
	Quality   string     `json:"quality"`
	Wendu     string     `json:"wendu"`
	Ganmao    string     `json:"ganmao"`
	Yesterday Yesterday  `json:"yesterday" xorm:"-"`
	Forecast  []Forecast `json:"forecast" xorm:"-"`
}
type Yesterday struct {
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float64 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}
type Forecast struct {
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float64 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}

type NewsData struct {
	Code   string    `json:"code"`
	Charge bool      `json:"charge"`
	Msg    string    `json:"msg"`
	Result NewsResult `json:"result"`
}
type NewsResult struct {
	Msg string `json:"msg"`
	ReS []string `json:"result"`
	Status string`json:"status"`
}

type NewsListData struct {
	Code   string    `json:"code"`
	Charge bool      `json:"charge"`
	Msg    string    `json:"msg"`
	Result NewsListResult `json:"result"`
}
type NewsListResult struct {
	Msg string `json:"msg"`
	ReS NewsList `json:"result"`
	Status string`json:"status"`
}
type NewsList struct {
	Num string `json:"num"`
	Channel string `json:"channel"`
	NewList []NewList `json:"list"`
}
type NewList struct {
	Src string `json:"src"`
	WebUrl string `json:"weburl"`
	Time string `json:"time"`
	Pic string `json:"pic"`
	Title string `json:"title"`
	Category string `json:"category"`
	Content string `json:"content"`
	Url string `json:"url"`
}

type AccData struct {
	Code   string    `json:"code"`
	Charge bool      `json:"charge"`
	Msg    string    `json:"msg"`
	Result AccResult `json:"result"`
}
type AccResult struct {
	Msg string `json:"msg"`
	ReS ReList `json:"result"`
}
type ReList struct {
	Num  string     `json:"num"`
	List []FoodList `json:"list"`
}
type FoodList struct {
	Id          int
	ClassId     string         `json:"classid"`
	Process     []ProcessList  `json:"process" xorm:"-"`
	PrepareTime string         `json:"preparetime"`
	Material    []MaterialList `json:"material" xorm:"-"`
	Name        string         `json:"name"`
	SId         string         `json:"id"`
	Pic         string         `json:"pic"`
	Tag         string         `json:"tag"`
	PeopleNum   string         `json:"peoplenum"`
	Content     string         `json:"content"`
	CookingTime string         `json:"cookingtime"`
	Status      string         `json:"status"`
}
type ProcessList struct {
	PContent string `json:"pcontent"`
	Pic      string `json:"pic"`
}
type MaterialList struct {
	Amount string `json:"amount"`
	MName  string `json:"mname"`
	Type   string `json:"type"`
}

type PhoneData struct {
	Code   string `json:"code"`
	Charge bool   `json:"charge"`
	Msg    string `json:"msg"`
	Result Phone  `json:"result"`
}
type Phone struct {
	AreaCode string `json:"areaCode"`
	Province string `json:"province"`
	City     string `json:"city"`
	CardType string `json:"cardType"`
	PostCode string `json:"postCode"`
	Operator string `json:"operator"`
	Status   int    `json:"status"`
}
type PhoneDataRe struct {
	Code   string  `json:"code"`
	Charge string  `json:"charge"`
	Msg    string  `json:"msg"`
	Result PhoneRe `json:"result"`
}
type PhoneRe struct {
	Msg    string   `json:"msg"`
	ReS    PhoneMap `json:"result"`
	Status string   `json:"status"`
}
type PhoneMap struct {
	ShouJi   string `json:"shouji"`
	Province string `json:"province"`
	City     string `json:"city"`
	Company  string `json:"company"`
	CardType string `json:"cardtype"`
	AreaCode string `json:"areacode"`
}

type Param struct {
	com_type string
	post_id  string
	city     string
	keyword  string
	phone    string
	shouji   string
	change   int
	news     string
	news_s   string
}

func MFrm() {
	fmt.Println(`
	----------
	主功能列表：
		1. 快递查询
		2. 天气查询
		3. 菜谱查询
		4. 手机号码归属地查询
		5. 新闻获取
		6. 退出

	说明：
		输入中文出现问题，请输入数字或英文。2017-09-08
		尽量请不要访问 2 和 3 接口，谢谢合作！2017-09-11
		天气接口已更新，可正常使用。3 号接口不要访问。2017-09-13
	----------
	`)
	num := ""
	for num == "" {
		num = Input("输入想要的功能：(功能列表数字)", "")
	}
	switch num {
	case "1":
		PFrm("快递查询", num)
	case "2":
		PFrm("天气查询", num)
	case "3":
		PFrm("菜谱查询", num)
	case "4":
		PFrm("手机号码归属地查询", num)
	case "5":
		PFrm("新闻查询", num)
	case "6":
		fmt.Println("退出系统")
		os.Exit(1)
	case "exit":
		fmt.Println("退出系统")
		os.Exit(1)
	default:
		fmt.Println("输入错误，请重新输入！")
		MFrm()
	}
}

func PFrm(str, num string) {
	p := Param{}
	fmt.Println("你选择的功能是：", str)
	switch num {
	case "1":
		com_type, post_id := ScanExpress()
		p = Param{com_type: com_type, post_id: post_id}
	case "2":
		city := ScanWeather()
		p = Param{city: city}
	case "3":
		keyword := ScanFood()
		p = Param{keyword: keyword}
	case "4":
		phone := ScanPhone()
		p = Param{phone: phone, shouji: phone}
	case "5":
		p = Param{news:"new"}
	}
	body, s, _ := GetBody(p)
	UnmarJson(body, num, s)
	MFrm()
}

func ScanExpress() (string, string) {
	fmt.Println(`
	----------
	快递公司编码：
		申通="shentong" EMS="ems" 顺丰="shunfeng"
		圆通="yuantong" 中通="zhongtong" 韵达="yunda"
		天天="tiantian" 汇通="huitongkuaidi" 全峰="quanfengkuaidi"
		德邦="debangwuliu" 宅急送="zhaijisong"
	----------
	`)
	com_type := ""
	post_id := ""
	for {
		com_type = Input("输入快递公司：", "")
		post_id = Input("输入快递单号：", "")
		if com_type != "" && post_id != "" {
			break
		}
	}
	//fmt.Printf("快递公司为：%v，单号为：%v。\n", com_type, post_id)
	return com_type, post_id
}

func PTypeOfNews(res []string) map[int]string {
	r := make(map[int]string)
	c := 1
	fmt.Printf("----------\n\t请选择新闻类型：\n")
	for k := 0; k < len(res); k++ {
		if k == 4 * c {
			fmt.Printf("\t%v. %v\t\n", k, res[k])
			c++
		} else {
			fmt.Printf("\t%v. %v\t", k, res[k])
		}
		r[k] = res[k]
	}
	fmt.Printf("\n----------")
	return r
}

func ScanNews(res map[int]string) string {
	fmt.Println(`
	----------
	新闻类型：
		输入功能列表序号，或详细类型名称如：头条
	----------
	`)
	n := ""
	for {
		n = Input("请输入：", "")
		if n != "" {
			break
		}
	}
	nn := ""
	for i := 0; i < len(res); i ++ {
		if n == strconv.Itoa(i) {
			nn = res[i]
			break
		} else if n == res[i] {
			nn = res[i]
			break
		} else {
			nn = ""
		}
	}
	if nn == "" {
		fmt.Println("输入错误，请重新输入！")
		return ScanNews(res)
	}
	return nn
}

func ScanFood() string {
	fmt.Println(`
	----------
	城市：
		比如：猪肉/白菜
	----------
	`)
	keyword := ""
	for {
		keyword = Input("输入搜索关键词：", "")
		if keyword != "" {
			break
		}
	}
	return keyword
}

func ScanPhone() string {
	fmt.Println(`
	----------
	城市：
		比如：188****2340
	----------
	`)
	phone := ""
	for {
		phone = Input("输入手机号码：", "")
		if phone != "" {
			break
		}
	}
	return phone
}

func ScanWeather() string {
	fmt.Println(`
	----------
	城市：
		比如：foshan/jiangmen
	----------
	`)
	city := ""
	for {
		city = Input("输入城市名称：", "")
		if city != "" {
			break
		}
	}
	//fmt.Printf("城市为：%v。\n", city)
	return city
}

func Input(say, defaults string) string {
	fmt.Println(say)
	var str string
	fmt.Scanln(&str)
	if strings.TrimSpace(str) == "" {
		if strings.TrimSpace(defaults) != "" {
			return defaults
		} else {
			fmt.Println("不能为空！")
			return Input(say, defaults)
		}
	}
	//fmt.Println("--" + str + "--")
	return str
}

func GetBody(p Param) ([]byte, string, error) {
	url := ""
	if p.city != "" {
		//url = "http://www.sojson.com/open/api/weather/json.shtml?city=" + strings.TrimSpace(p.city)
		url = "https://way.jd.com/he/freeweather?city=" + p.city + "&appkey=" + acc_key
	} else if p.com_type != "" && p.post_id != "" {
		url = "http://www.kuaidi100.com/query?type=" + strings.TrimSpace(p.com_type) + "&postid=" + strings.TrimSpace(p.post_id)
	} else if p.keyword != "" {
		url = "https://way.jd.com/jisuapi/search?keyword=" + p.keyword + "&num=" + strconv.Itoa(num) + "&appkey=" + acc_key
	} else if p.phone != "" && p.shouji != "" {
		url = "https://way.jd.com/shujujia/mobile?mobile=" + p.phone + "&appkey=" + acc_key
	} else if p.shouji != "" {
		url = "https://way.jd.com/jisuapi/query4?shouji=" + p.shouji + "&appkey=" + acc_key
	} else if p.news != "" {
		url = "https://way.jd.com/jisuapi/channel?appkey=" + acc_key
	} else if p.news_s != "" {
		url = "https://way.jd.com/jisuapi/get?channel="+p.news_s+"&num="+strconv.Itoa(num)+"&start="+start+"&appkey=" + acc_key
		sInt, _ := strconv.Atoi(start)
		sInt += 5
		start = strconv.Itoa(sInt)
	}
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败，请检查网络是否正常！")
		return nil, "", err
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取接口返回信息出错！")
		return nil, "", err
	}
	r.Body.Close()
	return body, p.shouji, nil
}

func UnmarJson(body []byte, num, str string) {
	switch num {
	case "1":
		logFile := WriteInit()
		logger := log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)
		defer logFile.Close()
		deli := new(Deli)
		err := json.Unmarshal(body, deli)
		if err != nil {
			fmt.Println("接口返回信息解析出错!")
		}
		exps := deli.Data
		for i := len(exps) - 1; i >= 0; i-- {
			fmt.Println("时间 ---> :", exps[i].TimeStr)
			fmt.Println("信息 ---> :", exps[i].Context)
			logger.Println("时间 ---> :", exps[i].TimeStr)
			logger.Println("信息 ---> :", exps[i].Context)
		}

	case "2":
		/*reback := new(ReBack)
		err := json.Unmarshal(body, reback)
		if err != nil {
			fmt.Println("接口返回信息解析出错!")
		}
		city := reback.City
		yes := reback.Data.Yesterday
		da := reback.Data
		das := reback.Data.Forecast
		fmt.Printf("--- > %v整体天气如下 \n", city)
		fmt.Printf("--- > 湿度 ：%v \n", da.Shidu)
		fmt.Printf("--- > 温度 ：%v \n", da.Wendu)
		fmt.Printf("--- > PM2.5：%v \n", da.Pm25)
		fmt.Printf("--- > PM10 ：%v \n", da.Pm10)
		fmt.Printf("--- > 空气质量：%v \n", da.Quality)
		fmt.Printf("--- > 感冒情况：%v \n", da.Ganmao)
		fmt.Println("")
		fmt.Printf("--- > %v未来五天天气如下 \n", city)
		for _, fcas := range das {
			fmt.Printf("--- > 日期：%v \n", fcas.Date)
			fmt.Printf("--- > 低温：%v \n", fcas.Low)
			fmt.Printf("--- > 高温：%v \n", fcas.High)
			fmt.Printf("--- > Aqi：%v \n", fcas.Aqi)
			fmt.Printf("--- > 风向：%v \n", fcas.Fx)
			fmt.Printf("--- > 风级：%v \n", fcas.Fl)
			fmt.Printf("--- > 天气：%v \n", fcas.Type)
			fmt.Printf("--- > 提醒：%v \n", fcas.Notice)
			fmt.Println("")
		}*/
		logFile := WriteInit()
		logger := log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)
		defer logFile.Close()
		weback := new(WeBack)
		err := json.Unmarshal(body, weback)
		if err != nil {
			fmt.Println("接口返回信息解析出错!")
		}
		if weback.Code != "10000" {
			ErrPrint(weback.Code)
		} else {
			nows := weback.Result.HeWeather5
			for _, n := range nows {
				fmt.Printf("--- > %v现在天气如下 \n", n.Basic.City)
				fmt.Printf("--- > 温度：%v \n", n.Now.Tmp)
				fmt.Printf("--- > 状态：%v \n", n.Now.Cond.Txt)
				fmt.Printf("--- > 风向：%v \n", n.Now.Wind.Dir)
				fmt.Printf("--- > 风级：%v \n", n.Now.Wind.Sc)
				fmt.Println("=== === === === === === ===")
				logger.Printf("--- > %v现在天气如下 \n", n.Basic.City)
				logger.Printf("--- > 温度：%v \n", n.Now.Tmp)
				logger.Printf("--- > 状态：%v \n", n.Now.Cond.Txt)
				logger.Printf("--- > 风向：%v \n", n.Now.Wind.Dir)
				logger.Printf("--- > 风级：%v \n", n.Now.Wind.Sc)
				logger.Println("=== === === === === === ===")
				sug := n.Suggestion
				fmt.Printf("--- > 天气事件建议 \n")
				fmt.Printf("--- > 外出：%v \n", sug.Uv.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Uv.Brf)
				fmt.Printf("--- > 洗车：%v \n", sug.Cw.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Cw.Brf)
				fmt.Printf("--- > 旅游：%v \n", sug.Trav.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Trav.Brf)
				fmt.Printf("--- > 空气质量：%v \n", sug.Air.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Air.Brf)
				fmt.Printf("--- > 舒适程度：%v \n", sug.Comf.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Comf.Brf)
				fmt.Printf("--- > 着装：%v \n", sug.Drsg.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Drsg.Brf)
				fmt.Printf("--- > 运动：%v \n", sug.Sport.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Sport.Brf)
				fmt.Printf("--- > 感冒：%v \n", sug.Flu.Txt)
				fmt.Printf("--- > 程度：%v \n", sug.Flu.Brf)
				fmt.Println("=== === === === === === ===")
				logger.Printf("--- > 天气事件建议 \n")
				logger.Printf("--- > 外出：%v \n", sug.Uv.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Uv.Brf)
				logger.Printf("--- > 洗车：%v \n", sug.Cw.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Cw.Brf)
				logger.Printf("--- > 旅游：%v \n", sug.Trav.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Trav.Brf)
				logger.Printf("--- > 空气质量：%v \n", sug.Air.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Air.Brf)
				logger.Printf("--- > 舒适程度：%v \n", sug.Comf.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Comf.Brf)
				logger.Printf("--- > 着装：%v \n", sug.Drsg.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Drsg.Brf)
				logger.Printf("--- > 运动：%v \n", sug.Sport.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Sport.Brf)
				logger.Printf("--- > 感冒：%v \n", sug.Flu.Txt)
				logger.Printf("--- > 程度：%v \n", sug.Flu.Brf)
				logger.Println("=== === === === === === ===")
				dailys := n.DailyForecast
				for _, d := range dailys {
					fmt.Printf("--- > %v 天气如下 \n", d.Date)
					fmt.Printf("--- > 最低温度：%v \n", d.Tmp.Min)
					fmt.Printf("--- > 最高温度：%v \n", d.Tmp.Max)
					fmt.Printf("--- > 天气状况：%v - %v \n", d.Cond.Txtn, d.Cond.Txtd)
					fmt.Printf("--- > 刮风情况：%v - %v \n", d.Wind.Sc, d.Wind.Dir)
					fmt.Println("--- --- --- ---")
					logger.Printf("--- > %v 天气如下 \n", d.Date)
					logger.Printf("--- > 最低温度：%v \n", d.Tmp.Min)
					logger.Printf("--- > 最高温度：%v \n", d.Tmp.Max)
					logger.Printf("--- > 天气状况：%v - %v \n", d.Cond.Txtn, d.Cond.Txtd)
					logger.Printf("--- > 刮风情况：%v - %v \n", d.Wind.Sc, d.Wind.Dir)
					logger.Println("--- --- --- ---")
				}
				fmt.Println("=== === === === === === ===")
				hours := n.HourlyForecast
				for _, h := range hours {
					fmt.Printf("--- > %v 天气如下 \n", h.Date)
					fmt.Printf("--- > 温度：%v \n", h.Tmp)
					fmt.Printf("--- > 天气状况：%v \n", h.Cond.Txt)
					fmt.Printf("--- > 刮风情况：%v - %v \n", h.Wind.Sc, h.Wind.Dir)
					fmt.Println("--- --- --- ---")
					logger.Printf("--- > %v 天气如下 \n", h.Date)
					logger.Printf("--- > 温度：%v \n", h.Tmp)
					logger.Printf("--- > 天气状况：%v \n", h.Cond.Txt)
					logger.Printf("--- > 刮风情况：%v - %v \n", h.Wind.Sc, h.Wind.Dir)
					logger.Println("--- --- --- ---")
				}
			}
		}

	case "3":
		logFile := WriteInit()
		logger := log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)
		defer logFile.Close()
		acc := new(AccData)
		err := json.Unmarshal(body, acc)
		if err != nil {
			fmt.Println("接口返回信息解析出错!")
		}
		if acc.Code != "10000" {
			ErrPrint(acc.Code)
		} else {
			lists := acc.Result.ReS.List
			fmt.Printf("--- > 菜谱如下 \n")
			logger.Printf("--- > 菜谱如下 \n")
			for _, list := range lists {
				fmt.Printf("--- > 菜名：%v \n", list.Name)
				fmt.Printf("--- > 标签：%v \n", list.Tag)
				fmt.Printf("--- > 食用人数：%v \n", list.PeopleNum)
				logger.Printf("--- > 菜名：%v \n", list.Name)
				logger.Printf("--- > 标签：%v \n", list.Tag)
				logger.Printf("--- > 食用人数：%v \n", list.PeopleNum)
				if strings.Contains(list.Content, "<br />") {
					list.Content = strings.Replace(list.Content, "<br />", "\n", -1)
				}
				fmt.Printf("--- > 内容：%v \n", list.Content)
				fmt.Printf("--- > 完成时间：%v \n", list.CookingTime)
				fmt.Println("")
				fmt.Printf("\t--- > 材料如下 \n")
				logger.Printf("--- > 内容：%v \n", list.Content)
				logger.Printf("--- > 完成时间：%v \n", list.CookingTime)
				logger.Println("")
				logger.Printf("\t--- > 材料如下 \n")
				mats := list.Material
				nName := ""
				amount := ""
				for _, mat := range mats {
					nName += mat.MName + "," + "---"
					amount += mat.Amount + "," + "---"
				}
				fmt.Printf("\t--- > 食料名称：%v \n", nName)
				fmt.Printf("\t--- > 食料用量：%v \n", amount)
				fmt.Println("")
				fmt.Printf("\t--- > 步骤如下 \n")
				logger.Printf("\t--- > 食料名称：%v \n", nName)
				logger.Printf("\t--- > 食料用量：%v \n", amount)
				logger.Println("")
				logger.Printf("\t--- > 步骤如下 \n")
				pros := list.Process
				for i := 0; i < len(pros); i++ {
					fmt.Printf("\t--- > %v. %v \n", i, pros[i].PContent)
				}
				fmt.Println("=== === === === === === ===")
				logger.Println("=== === === === === === ===")
			}
		}
	case "4":
		logFile := WriteInit()
		logger := log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)
		defer logFile.Close()
		pd := new(PhoneData)
		err := json.Unmarshal(body, pd)
		if err != nil {
			fmt.Println("接口返回信息解析出错!")
		}
		p := pd.Result
		if pd.Code != "10000" {
			ErrPrint(pd.Code)
			time.Sleep(time.Millisecond * 1200)
			fmt.Println("--- > 更换接口中... 请稍候! ")
			b, _, _ := GetBody(Param{shouji: str})
			re := new(PhoneDataRe)
			err := json.Unmarshal(b, re)
			if err != nil {
				fmt.Println("接口返回信息解析出错!")
			}
			if re.Code != "10000" {
				ErrPrint(pd.Code)
			} else {
				pMap := re.Result.ReS
				fmt.Printf("--- > 归属地信息如下 \n")
				fmt.Printf("--- > 手机号码： %v \n", pMap.ShouJi)
				fmt.Printf("--- > 手机号码公司： %v \n", pMap.Company)
				fmt.Printf("--- > 手机卡类型： %v \n", pMap.CardType)
				fmt.Printf("--- > 省份： %v \n", pMap.Province)
				fmt.Printf("--- > 城市： %v \n", pMap.City)
				fmt.Printf("\t--- > 区号： %v \n", pMap.AreaCode)
				logger.Printf("--- > 归属地信息如下 \n")
				logger.Printf("--- > 手机号码： %v \n", pMap.ShouJi)
				logger.Printf("--- > 手机号码公司： %v \n", pMap.Company)
				logger.Printf("--- > 手机卡类型： %v \n", pMap.CardType)
				logger.Printf("--- > 省份： %v \n", pMap.Province)
				logger.Printf("--- > 城市： %v \n", pMap.City)
				logger.Printf("\t--- > 区号： %v \n", pMap.AreaCode)
			}
		} else {
			fmt.Printf("--- > 归属地信息如下 \n")
			fmt.Printf("--- > 手机号码： %v \n", str)
			fmt.Printf("--- > 手机号码服务商： %v \n", p.Operator)
			fmt.Printf("--- > 手机卡类型： %v \n", p.CardType)
			fmt.Printf("--- > 省份： %v \n", p.Province)
			fmt.Printf("--- > 城市： %v \n", p.City)
			fmt.Printf("\t--- > 区号： %v \n", p.AreaCode)
			fmt.Printf("\t--- > 邮箱： %v \n", p.PostCode)
			logger.Printf("--- > 归属地信息如下 \n")
			logger.Printf("--- > 手机号码： %v \n", str)
			logger.Printf("--- > 手机号码服务商： %v \n", p.Operator)
			logger.Printf("--- > 手机卡类型： %v \n", p.CardType)
			logger.Printf("--- > 省份： %v \n", p.Province)
			logger.Printf("--- > 城市： %v \n", p.City)
			logger.Printf("\t--- > 区号： %v \n", p.AreaCode)
			logger.Printf("\t--- > 邮箱： %v \n", p.PostCode)
		}
	case "5":
		logFile := WriteInit()
		logger := log.New(logFile,"\r\n", log.Ldate | log.Ltime | log.Lshortfile)
		defer logFile.Close()
		n := new(NewsData)
		err := json.Unmarshal(body, n)
		if err != nil {
			fmt.Println("接口返回信息解析出错!")
		}
		if n.Code != "10000" {
			ErrPrint(n.Code)
		} else {
			res := n.Result.ReS
			ss := PTypeOfNews(res)
			s := ScanNews(ss)

			b, _, _ := GetBody(Param{news_s:s})
			nlist := new(NewsListData)
			err1 := json.Unmarshal(b, nlist)
			if err1 != nil {
				fmt.Println("接口返回信息解析出错!")
			}
			if nlist.Code != "10000" {
				ErrPrint(nlist.Code)
			} else {
				ns := nlist.Result.ReS.NewList
				fmt.Printf("--- > 你选择的新闻类型为：%v \n",nlist.Result.ReS.Channel)
				for _, vs := range ns {
					fmt.Println("=== === === === === === ===")
					fmt.Printf("\t--- > 标题： %v \n", vs.Title)
					fmt.Printf("\t--- > 时间： %v \t\t来源：%v \n", vs.Time, vs.Src)
					fmt.Printf("\t--- > 内容： %v \n", vs.Content)
					fmt.Println("=== === === === === === ===")
					fmt.Println("")
					logger.Println("=== === === === === === ===")
					logger.Printf("\t--- > 标题： %v \n", vs.Title)
					logger.Printf("\t--- > 时间： %v \t\t，来源：%v \n", vs.Time, vs.Src)
					logger.Printf("\t--- > 内容： %v \n", vs.Content)
					logger.Println("=== === === === === === ===")
					logger.Println("")
				}
			}

		}
	}
}

func ErrPrint(code string) {
	switch code {
	case "10001":
		fmt.Println("错误的请求appkey")
	case "11010":
		fmt.Println("接口调用异常，请稍后再试")
	case "11030":
		fmt.Println("接口返回格式有误")
	case "10003":
		fmt.Println("不存在相应的数据信息")
	case "10004":
		fmt.Println("URL上appkey参数不能为空")
	case "10010":
		fmt.Println("接口需付费")
	case "10020":
		fmt.Println("	系统繁忙，请稍后再试	")
	case "10030":
		fmt.Println("调用网关失败， 请与作者 lightsylvanus@foxmail.com联系")
	case "10040":
		fmt.Println("超过每天限量，请明天继续")
	case "10050":
		fmt.Println("用户已被禁用")
	case "10060":
		fmt.Println("	提供方设置调用权限，请联系作者 lightsylvanus@foxmail.com")
	case "10070":
		fmt.Println("该数据只允许企业用户调用")
	case "10090":
		fmt.Println("文件大小超限，请上传小于1M的文件")
	default:
		fmt.Println("未知错误，请联系作者 lightsylvanus@foxmail.com")
	}
}

// 获取当前执行命令所在的位置
func GetCurrentPath() (string, error) {
	return os.Getwd()
}

func getCurrentTime() (string) {
	return time.Now().Format("20060102150405")
}

func WriteInit() *os.File {
	dir, _ := GetCurrentPath()
	suffix := getCurrentTime()
	dir = filepath.Join(dir, "out")
	err := MakeDir(dir)
	if err != nil {
		fmt.Println("创建文件夹出错!")
	}
	filename := "out/" + getFileName() + suffix + ".log"
	/*if checkFileIsExist(filename) {
		os.Create(filename)
	}*/
	logFile, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开log文件出错！")
	}
	return logFile
}

func OpenF(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR, 0540)
}

func MakeDir(filedir string) error {
	return os.MkdirAll(filedir, 0777)
}

func checkFileIsExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func getFileName() string {
	_, fulleFilename, _, _ := runtime.Caller(0)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fulleFilename)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)

	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

func main() {
	MFrm()
}
