package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Id          int    `gorm:"primaryKey" json:"gid" label:"货物编号"`
	GoodsName   string `gorm:"type:varchar(20);not null" json:"goodsname" validate:"min=2,max=12" label:"货物名称"`
	MerchantId  int    `gorm:"type:int;not null" json:"mechantid"`
	Stock       int    `gorm:"type:int;not null" json:"stock" label:"货物库存"`
	Description string `gorm:"type:varchar(200)" json:"description" label:"货物描述"`
	State       int    `gorm:"type:int;default:0" json:"gstate" validate:"required,gte=2" label:"货物状态"`
}
