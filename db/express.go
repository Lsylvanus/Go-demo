package db

import "time"

type Delivery struct {
	Id        int64     `xorm:"autoincr pk"`
	Message   string    `json:"message"`
	Nu        string    `json:"nu"`
	IsCheck   string    `json:"ischeck" xorm:"'ischeck'"`
	Condition string    `json:"condition"`
	Com       string    `json:"com"`
	Status    string    `json:"status"`
	State     string    `json:"state"`
	Data      []Express `json:"data" xorm:"-"`
	Updated   time.Time `xorm:"updated"`
	Created   time.Time `xorm:"created"`
	Version   int       `xorm:"version"`
}

type Express struct {
	Id       int64 `xorm:"autoincr pk"`
	Nu       string
	TimeStr  string `json:"time" xorm:"-"`
	Time     time.Time
	FTimeStr string    `json:"ftime" xorm:"-"`
	FTime    time.Time `xorm:"'ftime'"`
	Context  string    `json:"context"`
	Location string    `json:"location"`
	Updated  time.Time `xorm:"updated"`
	Created  time.Time `xorm:"created"`
	Version  int       `xorm:"version"`
}
