package model

import "time"

type Product struct {
	ID          uint      `json:"product_id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	CommentNum  int       `json:"comment_num"`
	Price       float64   `json:"price"`
	IsAddedCart bool      `json:"is_addedCart"`
	Cover       string    `json:"cover"`
	PublishTime time.Time `json:"publish_time"`
	Link        string    `json:"link"`
}
