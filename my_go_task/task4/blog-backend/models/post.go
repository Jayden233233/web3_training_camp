package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments"`
}