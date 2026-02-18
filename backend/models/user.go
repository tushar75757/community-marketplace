package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Name             string         `gorm:"size:255;not null" json:"name"`
	Email            string         `gorm:"size:255;unique;not null" json:"email"`
	PasswordHash     string         `gorm:"size:255;not null" json:"-"`
	Role             string         `gorm:"type:enum('admin', 'user');default:'user'" json:"role"`
	CreatedByLoginId uint           `json:"created_by_login_id"`
	CreatedByRoleId  string         `gorm:"size:50" json:"created_by_role_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
