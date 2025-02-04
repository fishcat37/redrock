package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"redrock/model"
)

var DBComment *gorm.DB

func InitComment(dsn string) error {
	var err error
	DBComment, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DBComment.AutoMigrate(&model.Product{})
	if err != nil {
		return err
	}
	return nil
}

func GetComment(comment model.Comment) ([]model.Comment, error) {
	var comments []model.Comment
	result := DBComment.Model(&model.Comment{}).Where("product_id = ?", comment.ProductID).Find(&comments)
	if result.RowsAffected == 0 {
		return []model.Comment{}, nil
	} else if result.Error != nil {
		return []model.Comment{}, result.Error
	}
	return comments, nil
}

func AddComment(comment *model.Comment) error {
	result := DBComment.Model(&model.Comment{}).Create(&comment)
	return result.Error
}

func DeleteComment(comment model.Comment) error {
	result := DBComment.Model(&model.Comment{}).Where("id = ?", comment.ID).Delete(&comment)
	return result.Error
}

func UpdateComment(comment *model.Comment) error {
	result := DBComment.Model(&model.Comment{}).Where("id = ?", comment.ID).Update("content", comment.Content)
	return result.Error
}

func Praise(praise model.Praise) error {
	if praise.Model == 1 {
		result := DBComment.Model(&model.Comment{}).Where("comment_id = ?", praise.CommentID).Update("praise_count", gorm.Expr("praise_count + 1"))
		return result.Error
	} else if praise.Model == 2 {
		var comment model.Comment
		result := DBComment.Model(&model.Comment{}).Where("comment_id = ?", praise.CommentID).First(&comment)
		if result.RowsAffected == 0 {
			return result.Error
		}
		if comment.PraiseCount != 0 {
			result = DBComment.Model(&model.Comment{}).Where("comment_id = ?", praise.CommentID).Update("praise_count", gorm.Expr("praise_count - 1"))
			return result.Error
		} else {
			return fmt.Errorf("点赞已为零，无法降低")
		}
	}
	return nil

}
