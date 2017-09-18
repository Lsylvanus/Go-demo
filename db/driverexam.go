package db

type ApiUser struct {
	Id         int
	User       string
	UserCode   string
	UserStatus int
	ApiSite    string
	AccessKey  string
	Desc       string
}

type DriverEx struct {
	Code   string `json:"code"`
	Charge bool   `json:"charge"`
	Msg    string `json:"msg"`
	Result Result `json:"result"`
}

type Result struct {
	Msg    string `json:"msg"`
	ExamRe ExamRe `json:"result"`
}

type ExamRe struct {
	Total    string     `json:"total"`
	Subject  string     `json:"subject"`
	PageSize string     `json:"pagesize"`
	Sort     string     `json:"sort"`
	PageNum  string     `json:"pagenum"`
	Type     string     `json:"type"`
	List     []ExamList `json:"list"`
	Status   string     `json:"status"`
}

type ExamList struct {
	Id       int
	Explain  string `json:"explain"`
	Subject  string
	Chapter  string `json:"chapter"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
	Option3  string `json:"option3"`
	Option4  string `json:"option4"`
	Pic      string `json:"pic"`
	Type     string `json:"type"`
}
