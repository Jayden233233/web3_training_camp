package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 连接到SQLite内存数据库
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移模型
	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}); err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}

	// 创建测试数据
	createTestData(db)

	// 执行查询示例
	queryUserPostsWithComments(db, 1)
	queryPostWithMostComments(db)
}

func createTestData(db *gorm.DB) {
	// 创建用户
	user1 := User{Name: "Alice", Email: "alice@example.com"}
	user2 := User{Name: "Bob", Email: "bob@example.com"}
	db.Create(&user1)
	db.Create(&user2)

	// 创建文章
	post1 := Post{
		Title:         "GORM入门指南",
		Content:       "这是一篇关于GORM的入门文章...",
		UserID:        user1.ID,
		Status:        "发布",
		CommentNum:    0,
		CommentStatus: "无评论",
	}
	post2 := Post{
		Title:         "SQLite高级技巧",
		Content:       "探索SQLite的高级功能和优化策略...",
		UserID:        user1.ID,
		Status:        "发布",
		CommentNum:    0,
		CommentStatus: "无评论",
	}
	post3 := Post{
		Title:         "Go语言并发编程",
		Content:       "了解Go语言中的goroutine和channel...",
		UserID:        user2.ID,
		Status:        "草稿",
		CommentNum:    0,
		CommentStatus: "无评论",
	}
	db.Create(&post1)
	db.Create(&post2)
	db.Create(&post3)

	// 创建评论
	comment1 := Comment{
		Content: "非常实用的文章！",
		UserID:  user2.ID,
		PostID:  post1.ID,
	}
	comment2 := Comment{
		Content: "期待更多示例",
		UserID:  user2.ID,
		PostID:  post1.ID,
	}
	comment3 := Comment{
		Content: "讲得很清楚",
		UserID:  user1.ID,
		PostID:  post2.ID,
	}
	db.Create(&comment1)
	db.Create(&comment2)
	db.Create(&comment3)
}

func queryUserPostsWithComments(db *gorm.DB, userID uint) {
	var user User
	db.Preload("Posts.Comments").First(&user, userID)

	log.Printf("用户 %s 的文章:", user.Name)
	for _, post := range user.Posts {
		log.Printf("- 文章: %s (评论数: %d, 状态: %s)", post.Title, post.CommentNum, post.CommentStatus)
		for _, comment := range post.Comments {
			var commenter User
			db.First(&commenter, comment.UserID)
			log.Printf("  - 评论: %s (作者: %s)", comment.Content, commenter.Name)
		}
	}
}

func queryPostWithMostComments(db *gorm.DB) {
	var post Post
	db.Order("comment_num DESC").First(&post)
	log.Printf("评论最多的文章: %s (评论数: %d)", post.Title, post.CommentNum)
}
