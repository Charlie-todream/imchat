package models

import "time"

const (
	COMMUNITY_CATE_COM = 1
)

type Community struct {
	Id int64 `gorm:"primary_key;not_null;auto_increment" form:"id" json:"id"`
	//名称
	Name string `gorm:"varchar(30)" form:"name" json:"name"`
	//群主ID
	Ownerid int64 `gorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 什么角色
	//群logo
	Icon string `gorm:"varchar(250)" form:"icon" json:"icon"`
	//como
	Cate int64 `gorm:"int(11)" form:"cate" json:"cate"` // 什么角色
	//描述
	Memo string `gorm:"varchar(120)" form:"memo" json:"memo"` // 什么角色
	//
	Createat time.Time `gorm:"datetime" form:"createat" json:"createat"` // 什么角色
}
