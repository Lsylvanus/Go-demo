package db

type GC struct {
	Lon      float64 `json:"lon"`
	Level    int     `json:"level"`
	Address  string  `json:"address"`
	CityName string  `json:"cityName"`
	ALevel   int     `json:"alevel"`
	Lat      float64 `json:"lat"`
}

type NearBuilding struct {
	Status string `json:"status"`
	Data   Loc    `json:"data"`
}

type Loc struct {
	Id             int64
	Province       string      `json:"province"`
	CrossList      []CrossList `json:"cross_list" xorm:"-"`
	Code           string      `json:"code"`
	Tel            string      `json:"tel"`
	CityAdCode     string      `json:"cityadcode"`
	AreaCode       string      `json:"areacode"`
	TimeStamp      string      `json:"timestamp"`
	SeaArea        SeaArea     `json:"sea_area" xorm:"-"`
	Pos            string      `json:"pos"`
	RoadList       []RoadList  `json:"road_list" xorm:"-"`
	Result         string      `json:"result"`
	Message        string      `json:"message"`
	Desc           string      `json:"desc"`
	City           string      `json:"city"`
	DistrictAdCode string      `json:"districtadcode"`
	District       string      `json:"district"`
	Country        string      `json:"country"`
	ProvinceAdCode string      `json:"provinceadcode"`
	Version        string      `json:"version"`
	AdCode         string      `json:"adcode"`
	PoiList        []PoiList   `json:"poi_list" xorm:"-"`
}

type SeaArea struct {
	Id     int64
	LocId  int64
	AdCode string `json:"adcode"`
	Name   string `json:"name"`
}

type CrossList struct {
	Id        int64
	LocId     int64
	Distance  string `json:"distance"`
	Direction string `json:"direction"`
	Name      string `json:"name"`
	Weight    string `json:"weight"`
	Level     string `json:"level"`
	Longitude string `json:"longitude"`
	CrossId   string `json:"crossid"`
	Width     string `json:"width"`
	Latitude  string `json:"latitude"`
}

type RoadList struct {
	Id        int64
	LocId     int64
	Distance  string `json:"distance"`
	Direction string `json:"direction"`
	Name      string `json:"name"`
	Level     string `json:"level"`
	Longitude string `json:"longitude"`
	Width     string `json:"width"`
	RoadId    string `json:"roadid"`
	Latitude  string `json:"latitude"`
}

type PoiList struct {
	Id        int64
	LocId     int64
	Distance  string      `json:"distance"`
	Direction string      `json:"direction"`
	Tel       string      `json:"tel"`
	Name      string      `json:"name"`
	Weight    string      `json:"weight"`
	TypeCode  string      `json:"typecode"`
	Longitude string      `json:"longitude"`
	Address   string      `json:"address"`
	Latitude  string      `json:"latitude"`
	Entrances []Entrances `json:"entrances" xorm:"-"`
	Type      string      `json:"type"`
	PoiId     string      `json:"poiid"`
}

type Entrances struct {
	Id        int64
	LocId     int64
	PoiListId int64
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}
