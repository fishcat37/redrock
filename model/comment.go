package model

import "time"

type Comment struct {
	ID          uint      `gorm:"primary_key" json:"post_id"`
	PublishTime time.Time `json:"publish_time" gorm:"type:date;default:CURRENT_DATE"`
	Content     string    `json:"content"`
	UserID      uint      `json:"user_id"`
	Avatar      string    `json:"avatar"`
	NickName    string    `json:"nickname"`
	PraiseCount int       `json:"praise_count"`
	IsPraised   int       `json:"is_praised"`
	ProductID   uint      `json:"product_id"`
}
