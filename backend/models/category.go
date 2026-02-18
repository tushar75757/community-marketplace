package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Name             string         `gorm:"size:255;unique;not null" json:"name"`
	CreatedByLoginId uint           `json:"created_by_login_id"`
	CreatedByRoleId  string         `gorm:"size:50" json:"created_by_role_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Items            []Item         `gorm:"constraint:OnDelete:CASCADE;" json:"items,omitempty"`
}
