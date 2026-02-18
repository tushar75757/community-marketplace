package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Title            string         `gorm:"size:255;not null;index:idx_items_title" json:"title"`
	Description      string         `gorm:"type:text" json:"description"`
	Price            float64        `gorm:"type:decimal(10,2);not null;index:idx_items_price" json:"price"`
	CategoryId       uint           `gorm:"not null;index:idx_items_category_id" json:"category_id"`
	Category         Category       `gorm:"foreignKey:CategoryId" json:"category,omitempty"`
	Condition        string         `gorm:"type:enum('new', 'used');not null" json:"condition"`
	ImageUrl         string         `gorm:"size:255" json:"image_url"`
	SellerId         uint           `gorm:"not null" json:"seller_id"`
	Seller           User           `gorm:"foreignKey:SellerId" json:"seller,omitempty"`
	Status           string         `gorm:"type:enum('active', 'inactive', 'sold');default:'active'" json:"status"`
	CreatedByLoginId uint           `json:"created_by_login_id"`
	CreatedByRoleId  string         `gorm:"size:50" json:"created_by_role_id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
