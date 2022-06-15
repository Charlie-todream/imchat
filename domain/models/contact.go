package models

import "time"

const (
	CONCAT_CATE_USER     = 1
	CONCAT_CATE_COMUNITY = 2
)

type Contact struct {
	Id int64 `gorm:"primary_key;not_null;auto_increment" form:"id" json:"id"`
	//谁的10000
	Ownerid int64 `gorm:"bigint(10)" form:"ownerid" json:"ownerid"` //是否在线
	//对端,10001
	Dstobj int64 `gorm:"bigint(20)" form:"dstobj" json:"dstobj"` // 对端信息
	//
	Cate int    `gorm:"int(11)" form:"cate" json:"cate"`      // 什么类型
	Memo string `gorm:"varchar(120)" form:"memo" json:"memo"` // 备注
	//
	Createat time.Time `gorm:"datetime" form:"createat" json:"createat"` // 创建时间
}
