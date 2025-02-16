package model

import "time"

type Order struct {
	ID       uint             `gorm:"primary_key" uri:"order_id" json:"id"`
	UserID   uint             `json:"user_id" uri:"user_id"`
	Products []OrderedProduct `json:"orders" gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Address  string           `json:"address"`
	Total    uint             `json:"total"`
}

type OrderedProduct struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	ProductID   uint      `json:"product_id"`
	OrderID     uint      `json:"order_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CommentNum  int       `json:"comment_num"`
	Price       float64   `json:"price"`
	Cover       string    `json:"cover"`
	PublishTime time.Time `json:"publish_time"`
	Link        string    `json:"link"`
	Count       int       `json:"count" gorm:"default:1"`
}

// type OrderItem struct {
// 	OrderID   uint `json:"order_id"`
// 	ProductID uint `json:"product_id"`
// }
