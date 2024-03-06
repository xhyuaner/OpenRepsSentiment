package models

import (
	"fmt"
	"time"
)

type Area struct {
	AreaId    int     `gorm:"primary_key"  json:"area_id"`
	Login     string  `gorm:"type:varchar(255)" json:"login"`
	Latitude  float64 `gorm:"type:double" json:"latitude"`
	Longitude float64 `gorm:"type:double" json:"longitude"`
	AreaName  string  `gorm:"type:varchar(255)" json:"area_name"`
}

// TableName 重写表名
func (Area) TableName() string {
	return "developarea"
}

type JsonTime time.Time

// MarshalJSON 自定义一种时间类型字段的Json序列化规则
func (j JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type Sentiment struct {
	Id         int      `gorm:"primary_key" json:"id"`
	StartDate  JsonTime `gorm:"type:date" json:"startdate"`
	EndDate    JsonTime `gorm:"type:date" json:"enddate"`
	MblogTotal int      `gorm:"type:int" json:"mblog_total"`
	PosNum     int      `gorm:"type:int" json:"pos_num"`
	NegNum     int      `gorm:"type:int" json:"neg_num"`
	TimeFlag   int      `gorm:"type:int;comment:'0-周新增 1-月新增 4-月累计 5-周累计'" json:"time_flag"`
	AreaId     int      `gorm:"type:int"  json:"area_id"`
	Area       *Area    `gorm:"foreign_key:AreaId;references:AreaId"  json:"area"`
	RepoName   string   `gorm:"type:varchar(150)" json:"repo_name"`
}

// TableName 重写表名
func (Sentiment) TableName() string {
	return "developsentiment"
}
