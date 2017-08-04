package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/thinkboy/log4go"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var gLogger log4go.Logger
var engine *xorm.Engine
var session *xorm.Session

//返回Json
type Retrieve struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    []Data `json:"data"` //cannot unmarshal array into Go value of type main.Order
	Paging  struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"paging"`
}

type Data struct {
	OrderWithDetail `json:"Order"` //cannot unmarshal object into Go value of type string
}

type OrderWithDetail struct {
	Order
	//因json中Order与ShippingDetail有相同字段state，如该结构体中存在ShippingDetail，则Order中的state获取不了
	//ShippingDetail
}

//订单信息
type Order struct {
	//Sid                          int            `xorm:"autoincr int(11) unique default 0 notnull"`
	OrderId                      string         `json:"order_id" xorm:"varchar(64) pk notnull"`
	TransactionId                string         `json:"transaction_id" xorm:"varchar(64) null"`
	ProductId                    string         `json:"product_id" xorm:"varchar(64) null"`
	VariantId                    string         `json:"variant_id" xorm:"varchar(64) null"`
	BuyerId                      string         `json:"buyer_id" xorm:"varchar(64) null"`
	Quantity                     string         `json:"quantity" xorm:"varchar(4) null"`
	Sku                          string         `json:"sku" xorm:"varchar(32) null"`
	Size                         string         `json:"size" xorm:"varchar(32) null"`
	Color                        string         `json:"color" xorm:"varchar(16) null"`
	State                        string         `json:"state" xorm:"varchar(32) null"`
	ShippingProvider             string         `json:"shipping_provider" xorm:"varchar(64) null"`
	TrackingNumber               string         `json:"tracking_number" xorm:"varchar(64) null"`
	ShippedDate                  string         `json:"shipped_date" xorm:"varchar(32) null"`
	ShipNote                     string         `json:"ship_note" xorm:"varchar(200) null"`
	LastUpdated                  string         `json:"last_updated" xorm:"varchar(32) null"`
	OrderTotal                   string         `json:"order_total" xorm:"varchar(8) null"`
	DaysToFulfill                string         `json:"days_to_fulfill" xorm:"varchar(8) null"`
	HoursToFulfill               string         `json:"hours_to_fulfill" xorm:"varchar(8) null"`
	ExpectedShipDate             string         `json:"expected_ship_date" xorm:"varchar(32) null"`
	Price                        string         `json:"price" xorm:"varchar(8) null"`
	Cost                         string         `json:"cost" xorm:"varchar(16) null"`
	Shipping                     string         `json:"shipping" xorm:"varchar(4) null"`
	ShippingCost                 string         `json:"shipping_cost" xorm:"varchar(16) null"`
	ProductName                  string         `json:"product_name" xorm:"varchar(200) null"`
	ProductImageUrl              string         `json:"product_image_url" xorm:"varchar(500) null"`
	OrderTime                    string         `json:"order_time" xorm:"varchar(32) null"`
	ShippingDetail               ShippingDetail `json:"ShippingDetail"`
	RefundedBy                   string         `json:"refunded_by" xorm:"varchar(32) null"`
	RefundedTime                 string         `json:"refunded_time" xorm:"varchar(32) null"`
	RefundedReason               string         `json:"refunded_reason" xorm:"varchar(200) null"`
	IsWishExpress                string         `json:"is_wish_express" xorm:"varchar(16) null"`
	RequiresDeliveryConfirmation string         `json:"requires_delivery_confirmation" xorm:"varchar(16) null"`
	WeRequiredDeliveryDate       string         `json:"we_required_delivery_date" xorm:"varchar(32) null"`
	TrackingConfirmed            string         `json:"tracking_confirmed" xorm:"varchar(16) null"`
	TrackingConfirmedDate        string         `json:"tracking_confirmed_date" xorm:"varchar(32) null"`
	Name                         string         `xorm:"varchar(32) null"`
	StreetAddress1               string         `xorm:"varchar(200) null"`
	StreetAddress2               string         `xorm:"varchar(200) null"`
	City                         string         `xorm:"varchar(32) null"`
	DetailState                  string         `xorm:"varchar(32) null"`
	Country                      string         `xorm:"varchar(32) null"`
	Zipcode                      string         `xorm:"varchar(8) null"`
	PhoneNumber                  string         `xorm:"varchar(16) null"`
}

//发货详细
type ShippingDetail struct {
	//Vid            int    `xorm:"autoincr int(11) pk notnull"`
	Name           string `json:"name" xorm:"varchar(32) null"`
	StreetAddress1 string `json:"street_address1" xorm:"varchar(200) null"`
	StreetAddress2 string `json:"street_address2" xorm:"varchar(200) null"`
	City           string `json:"city" xorm:"varchar(32) null"`
	State          string `json:"state" xorm:"varchar(32) null"`
	Country        string `json:"country" xorm:"varchar(32) null"`
	Zipcode        string `json:"zipcode" xorm:"varchar(8) pk notnull"`
	PhoneNumber    string `json:"phone_number" xorm:"varchar(16) null"`
}

//接口地址参数
type Uri struct {
	Access_token string
	Since        string
	Limit        string
	Start        string
}

//错误信息返回
type ErrorMsg struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    int    `json:"data"`
}

// 从 API 中获取订单信息(JSON)，并解析成 struct
func fetchOrder(retrieve *Retrieve, start string, limit string, count int) {
	var order_b []byte
	//var error_b []byte

	uri := new(Uri)
	//个人token
	uri.Access_token = "8ddac2fa42794124893947b9ff02e857"
	uri.Since = "2017-06-01"

	//翻页，返回json数组数量限制大小为50
	const page int = 50

	var url = "https://merchant.wish.com/api/v2/order/get-fulfill?start=" + start + "&limit=" + limit + "&since=" + uri.Since + "&access_token=" + uri.Access_token

	gLogger.Info("请求的URL为：\n", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	order_b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	gLogger.Info("获取页面body")

	retrieve = &Retrieve{}
	//errorMsg = &ErrorMsg{}
	err = json.Unmarshal(order_b, retrieve) // JSON to Struct
	if err != nil {
		panic(err.Error())
	}
	/*err = json.Unmarshal(error_b, errorMsg)
	if err != nil {
		panic(err.Error())
	}*/

	// create table
	t_order := new(Order)
	//查找表是否存在
	b, err := engine.IsTableExist(t_order)
	if !b {
		//创建表
		err := engine.CreateTables(t_order)
		if err != nil {
			gLogger.Debug("创建表失败！\n", err.Error())
		}
	} else if err != nil {
		gLogger.Debug("查找表失败！\n", err.Error())
	} else {
		// do nothing //
		// return
	}

	data := retrieve.Data
	//没有数据直接返回
	if len(data) == 0 {
		//关掉资源
		deinitAll()
		return
	}
	//fmt.Println(data)

	var order Order
	var ship_detail ShippingDetail

	//遍历数组的data
	for i := 0; i < len(data); i++ {
		//fmt.Println(data[i])
		order = data[i].OrderWithDetail.Order
		ship_detail = data[i].OrderWithDetail.Order.ShippingDetail

		gLogger.Info("订单数据： \n", order)
		gLogger.Info("发货详细： \n", ship_detail)
		gLogger.Debug("解析json格式，当前读取数为：%d ", i)
		//fmt.Printf("================\torder.State：%s\t ship.State：%s\t================\n", order.State, order.ShippingDetail.State)

		//保存json数据
		logJSONdata(order, ship_detail, i)

		//避免order的重复，先查再判断
		_, err := engine.Cols("order_id").Get(&order)
		//err2 := engine.In("order_id", []string{order.OrderId}).Find(&order)
		if err != nil {
			gLogger.Debug("已存在 \n", err.Error())
		}

		//orderid不重复，插入，否则不插入
		if order.OrderId != data[i].OrderWithDetail.Order.OrderId {
			gLogger.Debug("数据库已有该id：" + order.OrderId + "的订单 \n")
		} else {
			//两个结构体存进一个表t_order中
			order.PhoneNumber = ship_detail.PhoneNumber
			order.City = ship_detail.City
			order.DetailState = ship_detail.State
			order.Name = ship_detail.Name
			order.Country = ship_detail.Country
			order.Zipcode = ship_detail.Zipcode
			order.StreetAddress1 = ship_detail.StreetAddress1
			order.StreetAddress2 = ship_detail.StreetAddress2

			affected, err := engine.Insert(order)
			if err != nil {
				gLogger.Debug("插入表失败！\n", err.Error())
			}
			affecttedStr := strconv.FormatInt(affected, 10)
			gLogger.Info("成功插入t_order " + affecttedStr + " 条记录~！\n")
		}

	}

	gLogger.Debug("当前页面数组长度： %d 与默认限制大小 %s 是否相等：%t", len(data), limit, len(data) == page)
	//超过默认的50条数据，可能还有数据，则翻页
	if len(data) == page {
		gLogger.Debug("翻页查询。。。\n")
		start = strconv.Itoa(page*count + 1)
		//API提供的接口中最大的限制为500	Limit can range from 1 to 500 items and the default is 50
		if start == "501" {
			gLogger.Debug("已收集达最大LIMIT（500）限制")
			//超过限制返回，关掉资源
			deinitAll()
			return
		}
		limit = strconv.Itoa(page)
		count += 1
		//递归调用，起始位置数据不同
		fetchOrder(retrieve, start, limit, count)
	}

	//小于50限制，获取完成关掉资源
	deinitAll()
}

//初始化
func initAll() {
	gLogger = nil

	initLogger()
	conn()

	gLogger.Info("初始化。。。")
}

//关闭资源
func deinitAll() {
	err := session.Commit()
	if err != nil {
		return
	}
	defer session.Close()

	if gLogger != nil {
		return
	} else {
		gLogger = nil
		gLogger.Close()
	}
	gLogger.Info("关掉所有资源")
}

//初始化logger
func initLogger() {
	filenameOnly := getCurFileName()
	logFilename := filenameOnly + ".log"

	//控制台输出
	//gLogger.AddFilter("stdout", log4go.INFO, log4go.NewConsoleLogWriter())
	gLogger = log4go.NewDefaultLogger(log4go.INFO)

	//log文件输出
	if _, err := os.Stat(logFilename); err == nil {
		fmt.Printf("found old log file %s, now remove it\n", logFilename)
		gLogger.Debug("已有旧文件： %s, 重新创建\n", logFilename)
		os.Remove(logFilename)
	}
	gLogger.AddFilter("log", log4go.FINEST, log4go.NewFileLogWriter(logFilename, false))
	gLogger.Debug("当前时间 : %s", time.Now().Format("2006-01-02 15:04:05"))

	return
}

// 获取当前文件名称
func getCurFileName() string {
	_, fullFilename, _, _ := runtime.Caller(0)
	filenameWithSuffix := path.Base(fullFilename)
	fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}

//数据库初始化
func conn() {
	//创建Orm引擎
	var err error
	engine, err = xorm.NewEngine("mysql", "root:chenghai3c@/wish_catch?charset=utf8&parseTime=true")
	if err != nil {
		println(err.Error())
	}
	gLogger.Info("创建Orm引擎，准备连接数据库。。。")
	//测试连接
	errPing := engine.Ping()
	if errPing != nil {
		println(errPing.Error())
	}
	gLogger.Info("连接数据库成功!")
	// defer engine.Close()

	//当使用事务处理时，需要创建Session对象。在进行事物处理时，可以混用ORM方法和RAW方法
	session = engine.NewSession()

	// add Begin() before any action
	err1 := session.Begin()
	if err1 != nil {
		println(err1.Error())
	}
	gLogger.Info("开启事务。。。")

	//日志是一个接口，通过设置日志，可以显示SQL，警告以及错误等，默认的显示级别为INFO
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	//日志信息保存
	f, err := os.Create("catch_sql.log")
	if err != nil {
		println(err.Error())
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	gLogger.Info("创建log文件，保存执行的SQL语句。")

	//连接池
	//最大打开连接数
	engine.SetMaxOpenConns(100)
	//连接池的空闲数大小
	engine.SetMaxIdleConns(20)

	// 映射规则
	// 默认对应的表名就变成了 t_order 了，而之前默认的是 order
	tMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	engine.SetTableMapper(tMapper)

	//最后一个值得注意的是时区问题，默认xorm采用Local时区，所以默认调用的time.Now()会先被转换成对应的时区。
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

}

//保存json数据存进log
func logJSONdata(order Order, detail ShippingDetail, i int) {
	gLogger.Debug("=============================== Data Order [%d] ===============================", i)
	gLogger.Debug("Order.order_id\t\t\t\t\t\t\t=%s", order.OrderId)
	gLogger.Debug("Order.transaction_id\t\t\t\t\t=%s", order.TransactionId)
	gLogger.Debug("Order.product_id\t\t\t\t\t\t=%s", order.ProductId)
	gLogger.Debug("Order.variant_id\t\t\t\t\t\t=%s", order.VariantId)
	gLogger.Debug("Order.buyer_id\t\t\t\t\t\t\t=%s", order.BuyerId)
	gLogger.Debug("Order.quantity\t\t\t\t\t\t\t=%s", order.Quantity)
	gLogger.Debug("Order.sku\t\t\t\t\t\t\t\t=%s", order.Sku)
	gLogger.Debug("Order.size\t\t\t\t\t\t\t\t=%s", order.Size)
	gLogger.Debug("Order.color\t\t\t\t\t\t\t\t=%s", order.Color)
	gLogger.Debug("Order.state\t\t\t\t\t\t\t\t=%s", order.State)
	gLogger.Debug("Order.shipping_provider\t\t\t\t\t=%s", order.ShippingProvider)
	gLogger.Debug("Order.tracking_number\t\t\t\t\t=%s", order.TrackingNumber)
	gLogger.Debug("Order.shipped_date\t\t\t\t\t\t=%s", order.ShippedDate)
	gLogger.Debug("Order.ship_note\t\t\t\t\t\t\t=%s", order.ShipNote)
	gLogger.Debug("Order.last_updated\t\t\t\t\t\t=%s", order.LastUpdated)
	gLogger.Debug("Order.order_total\t\t\t\t\t\t=%s", order.OrderTotal)
	gLogger.Debug("Order.days_to_fulfill\t\t\t\t\t=%s", order.DaysToFulfill)
	gLogger.Debug("Order.hours_to_fulfill\t\t\t\t\t=%s", order.HoursToFulfill)
	gLogger.Debug("Order.expected_ship_date\t\t\t\t=%s", order.ExpectedShipDate)
	gLogger.Debug("Order.price\t\t\t\t\t\t\t\t=%s", order.Price)
	gLogger.Debug("Order.cost\t\t\t\t\t\t\t\t=%s", order.Cost)
	gLogger.Debug("Order.shipping\t\t\t\t\t\t\t=%s", order.Shipping)
	gLogger.Debug("Order.shipping_cost\t\t\t\t\t\t=%s", order.ShippingCost)
	gLogger.Debug("Order.product_name\t\t\t\t\t\t=%s", order.ProductName)
	gLogger.Debug("Order.product_image_url\t\t\t\t\t=%s", order.ProductImageUrl)
	gLogger.Debug("Order.order_time\t\t\t\t\t\t=%s", order.OrderTime)
	gLogger.Debug("Order.shipping_detail\t\t\t\t\t=%s", order.ShippingDetail)
	gLogger.Debug("Order.refunded_by\t\t\t\t\t\t=%s", order.RefundedBy)
	gLogger.Debug("Order.refunded_time\t\t\t\t\t\t=%s", order.RefundedTime)
	gLogger.Debug("Order.refunded_reason\t\t\t\t\t=%s", order.RefundedReason)
	gLogger.Debug("Order.is_wish_express\t\t\t\t\t=%s", order.IsWishExpress)
	gLogger.Debug("Order.requires_delivery_confirmation\t=%s", order.RequiresDeliveryConfirmation)
	gLogger.Debug("Order.we_required_delivery_date\t\t\t=%s", order.WeRequiredDeliveryDate)
	gLogger.Debug("Order.tracking_confirmed\t\t\t\t=%s", order.TrackingConfirmed)
	gLogger.Debug("Order.tracking_confirmed_date\t\t\t=%s", order.TrackingConfirmedDate)
	gLogger.Debug("=============================== Data End [%d] ===============================\n", i)

	gLogger.Debug("=============================== Data Ship [%d] ===============================", i)
	gLogger.Debug("Order.name\t\t\t\t\t\t\t\t=%s", detail.Name)
	gLogger.Debug("Order.street_address1\t\t\t\t\t=%s", detail.StreetAddress1)
	gLogger.Debug("Order.street_address2\t\t\t\t\t=%s", detail.StreetAddress2)
	gLogger.Debug("Order.city\t\t\t\t\t\t\t\t=%s", detail.City)
	gLogger.Debug("Order.detail_state\t\t\t\t\t\t=%s", detail.State)
	gLogger.Debug("Order.country\t\t\t\t\t\t\t=%s", detail.Country)
	gLogger.Debug("Order.zipcode\t\t\t\t\t\t\t=%s", detail.Zipcode)
	gLogger.Debug("Order.phone_number\t\t\t\t\t\t=%s", detail.PhoneNumber)
	gLogger.Debug("=============================== Data End [%d] ===============================\n", i)
}

//main method
func main() {
	initAll()

	fetchOrder(new(Retrieve), "0", "50", 1)

	//go语言中用log4go输出信息时有bug：只输出部分信息，甚至是无任何输出
	time.Sleep(100 * time.Millisecond)
}
