package models

import "time"

type Message struct {
	ID         uint      `gorm:"primaryKey"`
	ChatUserID string    `gorm:"column:chat_user_id;index"`
	Role       string    `gorm:"column:role;type:text;not null"`
	Content    string    `gorm:"column:message;type:text;not null"`
	Time       time.Time `gorm:"column:time;autoCreateTime"`
}
