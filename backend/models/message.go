package models

import (
	"time"
)

type Message struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	SenderId         uint      `gorm:"not null;index:idx_messages_sender_receiver" json:"sender_id"`
	Sender           User      `gorm:"foreignKey:SenderId" json:"sender,omitempty"`
	ReceiverId       uint      `gorm:"not null;index:idx_messages_sender_receiver" json:"receiver_id"`
	Receiver         User      `gorm:"foreignKey:ReceiverId" json:"receiver,omitempty"`
	ItemId           uint      `gorm:"not null" json:"item_id"`
	Item             Item      `gorm:"foreignKey:ItemId" json:"item,omitempty"`
	Content          string    `gorm:"type:text;not null" json:"content"`
	IsRead           bool      `gorm:"default:false" json:"is_read"`
	CreatedByLoginId uint      `json:"created_by_login_id"`
	CreatedByRoleId  string    `gorm:"size:50" json:"created_by_role_id"`
	CreatedAt        time.Time `json:"created_at"`
}
