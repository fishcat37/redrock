package model

import "time"

type Order struct {
	ID      uint     `gorm:"primary_key"`
	UserID  uint     `json:"user_id"`
	Orders  []Orders `json:"orders"`
	Address string   `json:"address"`
	Total   uint     `json:"total"`
}
type Orders struct {
	ProductID   uint      `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CommentNum  int       `json:"comment_num"`
	Price       float64   `json:"price"`
	IsAddedCart bool      `json:"is_addedCart"`
	Cover       string    `json:"cover"`
	PublishTime time.Time `json:"publish_time"`
	Link        string    `json:"link"`
	Count       int       `json:"count" gorm:"count;default:1"`
}
