package models

type Chat struct {
	UserID   string    `gorm:"column:user_id;primaryKey"`
	Messages []Message `gorm:"foreignKey:ChatUserID;references:UserID"`
	IsClient bool      `gorm:"column:is_client;default:false"`
}
