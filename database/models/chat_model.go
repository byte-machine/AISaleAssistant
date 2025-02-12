package models

import (
	"github.com/lib/pq"
)

type Chat struct {
	UserID   string         `gorm:"column:user_id;primaryKey"`
	Messages pq.StringArray `gorm:"column:messages;type:text[]"`
	IsClient bool           `gorm:"column:is_client;default:false"`
}
