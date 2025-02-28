package models

import "time"

type WaitingChat struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement"`
	ChatUserID  string    `gorm:"column:chat_user_id;index"`
	Since       time.Time `gorm:"column:since;autoCreateTime"`
	IsReminded  bool      `gorm:"column:is_reminded;default:false"`
	RemindCount uint      `gorm:"column:remind_count;default:0"`
}
