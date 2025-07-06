package main

import (
	"gorm.io/gorm"
)

// 创建文章时更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// 删除评论时更新文章的评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)

	status := "有评论"
	if count == 0 {
		status = "无评论"
	}

	return tx.Model(&Post{}).Where("id = ?", c.PostID).Updates(map[string]interface{}{
		"comment_num":    count,
		"comment_status": status,
	}).Error
}
