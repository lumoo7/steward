package models

import (
	"fmt"
	"gorm.io/gorm"
	"steward/system"
	"steward/system/base"
	"steward/system/constant"
	"steward/system/utils"
)

type User struct {
	gorm.Model
	Code     uint   `gorm:"type:int;unique;comment:用户编号" json:"code"`
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
	Code      uint   `json:"code"`
	Username  string `json:"username"`
	Gender    int    `json:"gender"`
	Password  string `json:"password"`
	Status    int    `json:"status"`
	Identity  int    `json:"identity"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserDto struct {
	Code         uint   `json:"code"`
	Username     string `json:"username"`
	Status       int    `json:"status"`
	StatusName   string `json:"statusName"`
	Identity     int    `json:"identity"`
	IdentityName string `json:"identityName"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

func Transfer2User(stu *UserStu) *User {
	return &User{
		Code:     stu.Code,
		Username: stu.Username,
		Gender:   stu.Gender,
		Password: stu.Password,
		Status:   stu.Status,
		Identity: stu.Identity,
		Phone:    stu.Phone,
		Email:    stu.Email,
	}
}

func Transfer2UserDto(user *User) *UserDto {
	return &UserDto{
		Code:         user.Code,
		Username:     user.Username,
		Status:       user.Status,
		StatusName:   statusName(user.Status),
		Identity:     user.Identity,
		IdentityName: identityName(user.Identity),
		Phone:        user.Phone,
		Email:        user.Email,
		CreatedAt:    utils.FormatTime(user.CreatedAt),
		UpdatedAt:    utils.FormatTime(user.UpdatedAt),
	}
}

// statusName 把 status 参数转换为 statusName
func statusName(s int) string {
	switch s {
	case 1:
		return constant.UserStatusAtHome
	case 2:
		return constant.UserStatusOuter
	default:
		return constant.UserStatusUnknown
	}
}

// identityName 把 identity 参数转换为 identityName
func identityName(s int) string {
	switch s {
	case 0:
		return constant.IdentityHost
	case 1:
		return constant.IdentityGuest
	case 2:
		return constant.IdentityStranger
	default:
		return constant.IdentityUnknown
	}
}

func init() {
	fmt.Println("\ninit user")
	system.RegisterModel(&User{})
}
