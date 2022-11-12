package models

import (
	"login-signup-api/config"
	"gorm.io/gorm"
)

type ManageAccess struct {
	gorm.Model
	AdminID int    `json:"admin_id" gorm: "foreignKey:UserId"` // foreign key
	UserID  int    `json:"user_id" gorm: "foreignKey:UserId"`  // foreign key
	Role    string `json:"dashboard_name" gorm:"size:255;not null"`
}

// save access
func (access *ManageAccess) SaveAccess() (*ManageAccess, error) {
	err := config.DB.Create(&access).Error
	if err != nil {
		return &ManageAccess{}, err
	}
	return access, nil
}
