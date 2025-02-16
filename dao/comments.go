package dao

import (
	"fmt"
	"gorm.io/gorm"
	"redrock/model"
	"time"
)

func GetComment(comment model.Comment) ([]model.Comment, error) {
	var comments []model.Comment
	result := DB.Model(&model.Comment{}).Where("product_id = ?", comment.ProductID).Find(&comments)
	if result.RowsAffected == 0 {
		return []model.Comment{}, nil
	} else if result.Error != nil {
		return []model.Comment{}, result.Error
	}
	return comments, nil
}

func AddComment(comment *model.Comment) error {
	comment.PublishTime = time.Now()
	err := FindProductById(comment.ProductID)
	if err != nil {
		return fmt.Errorf("product %v does not exist", comment.ProductID)
	}
	result := DB.Model(&model.Comment{}).Create(&comment)
	return result.Error
}

func DeleteComment(comment model.Comment) error {
	result := DB.Model(&model.Comment{}).Where("id = ? AND user_id = ?", comment.ID, comment.UserID).Delete(&comment)
	if result.RowsAffected == 0 {
		return fmt.Errorf(fmt.Sprintf("your comment %v does not exist", comment.ID))
	}
	return result.Error
}

func UpdateComment(comment *model.Comment) error {
	result := DB.Model(&model.Comment{}).Where("id = ? AND user_id = ?", comment.ID, comment.UserID).Update("content", comment.Content)
	if result.RowsAffected == 0 {
		return fmt.Errorf(fmt.Sprintf("your comment %v does not exist", comment.ID))
	}
	return result.Error
}

func Praise(praise model.Praise) error {
	if praise.Model == 1 {
		result := DB.Model(&model.Comment{}).Where("id = ?", praise.CommentID).Update("praise_count", gorm.Expr("praise_count + 1"))
		if result.RowsAffected == 0 {
			return fmt.Errorf(fmt.Sprintf("comment %v does not exist", praise.CommentID))
		}
		return result.Error
	} else if praise.Model == 2 {
		var comment model.Comment
		result := DB.Model(&model.Comment{}).Where("id = ?", praise.CommentID).First(&comment)
		if result.RowsAffected == 0 {
			return fmt.Errorf(fmt.Sprintf("comment %v does not exist", praise.CommentID))
		}
		if comment.PraiseCount != 0 {
			result = DB.Model(&model.Comment{}).Where("id = ?", praise.CommentID).Update("praise_count", gorm.Expr("praise_count - 1"))
			return result.Error
		} else {
			return fmt.Errorf("点赞已为零，无法降低")
		}
	}
	return nil

}
