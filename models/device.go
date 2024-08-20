package models

import (
	"gorm.io/gorm"
	"steward/system"
)

type Device struct {
	gorm.Model
	Name         string `gorm:"type:varchar(50);unique;comment:设备名称" json:"name"`
	SerialNumber string `gorm:"type:varchar(50);unique;comment:序列号" json:"serialNumber"`
	Location     string `gorm:"type:varchar(255);default null;comment:位置" json:"location"`
	Type         string `gorm:"type:varchar(20);default null;comment:类型" json:"type"`
	Manufacturer string `gorm:"type:varchar(50);default null;comment:厂商" json:"manufacturer"`
	Status       int    `gorm:"type:int;default 0;comment:状态" json:"status"`
}

func init() {
	system.RegisterModel(&Device{})
}
