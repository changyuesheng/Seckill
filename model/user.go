package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `gorm:"primaryKey" json:"uid"`                                             // 面试问题
	Username string `gorm:"type:varchar(20);NOT NULL" json:"username" validate:"min=4,max=12"` //大坑，不能有空格
	Password string `gorm:"type:varchar(200);NOT NULL" json:"password" validate:"min=6,max=20"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,oneof=0 1 2"`
}
