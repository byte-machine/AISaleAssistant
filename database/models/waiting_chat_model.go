package models

import "time"

type WaitingChat struct {
	ID         uint      `gorm:"column:id;primaryKey;autoIncrement"`
	ChatUserID string    `gorm:"column:chat_user_id;index"`
	Since      time.Time `gorm:"column:since;autoCreateTime"`
}
