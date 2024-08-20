package businese

import (
	"fmt"
	"gorm.io/gorm"
	"steward/models"
	"steward/system/db"
)

type service struct {
	DB *gorm.DB
}

func NewService() *service {
	return &service{
		DB: db.DB(),
	}
}

// AddUser 保存用户
func (s *service) AddUser(stu *models.UserStu) error {
	if err := s.DB.Create(models.Transfer2User(stu)).Error; err != nil {
		return fmt.Errorf("add user error: %w", err)
	}
	return nil
}

// DeleteUser 删除用户
func (s *service) DeleteUser(code uint) error {
	if err := s.DB.Where("code = ?", code).Delete(&models.User{}).Error; err != nil {
		return fmt.Errorf("delete user error: %w", err)
	}
	return nil
}

// FindUser 获取用户
func (s *service) FindUser(stu *models.UserStu) (*models.User, error) {
	var user = new(models.User)
	if err := s.where(stu).Find(user).Error; err != nil {
		return nil, fmt.Errorf("find user error: %w", err)
	}
	return user, nil
}

// PageListUser 分页查询
func (s *service) PageListUser(stu *models.UserStu) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64
	if err := s.where(stu).Find(&users).Count(&total).Offset((stu.PageIndex - 1) * stu.PageSize).Limit(stu.PageSize).Error; err != nil {
		return nil, total, fmt.Errorf("page list error: %w", err)
	}
	return users, total, nil
}

// UpdateUser 更新用户
func (s *service) UpdateUser(user *models.User) (*models.User, error) {
	if err := s.DB.Updates(user).Error; err != nil {
		return nil, fmt.Errorf("update user error: %w", err)
	}
	return user, nil
}

func (s *service) where(stu *models.UserStu) *gorm.DB {
	db := s.DB.Model(&models.User{})
	if stu.Id != 0 {
		db = db.Where("id = ?", stu.Id)
	}
	if len(stu.Username) != 0 {
		db = db.Where("name like %?%", stu.Username)
	}
	if len(stu.Email) != 0 {
		db = db.Where("email like %?%", stu.Email)
	}
	if len(stu.Phone) != 0 {
		db = db.Where("phone like %?%", stu.Phone)
	}
	return db
}
