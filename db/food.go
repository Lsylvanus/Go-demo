package db

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
	Id       int
	SId      string
	PId      int    // 顺序
	PContent string `json:"pcontent"`
	Pic      string `json:"pic"`
}

type MaterialList struct {
	Id     int
	SId    string
	Amount string `json:"amount"`
	MName  string `json:"mname"`
	Type   string `json:"type"`
}

type BccData struct {
	Code   string    `json:"code"`
	Charge bool      `json:"charge"`
	Msg    string    `json:"msg"`
	Result BccResult `json:"result"`
}

type BccResult struct {
	Msg string       `json:"msg"`
	ReS []ParentList `json:"result"`
}

type ParentList struct {
	Id      int
	Num     string      `json:"num"`
	ClassId string      `json:"classid"`
	Name    string      `json:"name"`
	List    []ChildList `json:"list" xorm:"-"`
}

type ChildList struct {
	Id       int
	ClassId  string `json:"classid"`
	Name     string `json:"name"`
	ParentId string `json:"parentid"`
}
