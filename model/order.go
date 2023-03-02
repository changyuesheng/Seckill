package model

import "gorm.io/gorm"

type Order struct {
	// Goods Goods `gorm:"foreignkey:Oid"`
	// User  User  `gorm:"foreignkey:Uid"`
	gorm.Model
	Id        int    `gorm:"primary_key;auto_increment" json:"oid" label:"订单编号"`
	GoodsName string `gorm:"type:varchar(20);not null" json:"goodsname" validate:"min=2,max=12" label:"货物名称"`
	Username  string `gorm:"type:varchar(20);NOT NULL" json:"username" validate:"min=4,max=12"` //大坑，不能有空格
	State     int    `gorm:"type:int;default:0" json:"ostate" validate:"required" label:"订单状态"`
}
