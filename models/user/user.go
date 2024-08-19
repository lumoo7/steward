package user

import (
	"gorm.io/gorm"
	"steward/common"
	"steward/common/base"
)

type User struct {
	gorm.Model
	Code     int    `gorm:"type:int;unique;comment:用户编号" json:"code"`
	Username string `gorm:"type:varchar(20);not null;comment:用户名" json:"username"`
	Gender   int    `gorm:"type:int;default:0;comment:性别（0女，1男）" json:"gender"`
	Password string `gorm:"type:varchar(50);not null;comment:密码" json:"password"`
	Status   int    `gorm:"type:int;default:0;comment:状态" json:"status"`
	Identity int    `gorm:"type:int;default:0;comment:身份" json:"identity"`
	Phone    string `gorm:"type:varchar(20);not null;comment:手机号码" json:"phone"`
	Email    string `gorm:"type:varchar(50);default:null;comment:电子邮箱" json:"email"`
}

type UserStu struct {
	base.Page
	base.Sort
	Id        uint   `json:"id"`
	Code      int    `json:"code"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Identity  int    `json:"identity"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserDto struct {
	Code         int    `json:"code"`
	Name         string `json:"name"`
	Status       int    `json:"status"`
	StatusName   string `json:"statusName"`
	Identity     int    `json:"identity"`
	IdentityName string `json:"identityName"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

func init() {
	common.RegisterModel(&User{})
}
