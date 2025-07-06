package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;size:50"`
	Email     string `gorm:"unique;not null;size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Posts     []Post         `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PostCount int            `gorm:"default:0"`
}

type Post struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"not null;size:100"`
	Content       string `gorm:"type:text"`
	UserID        uint   `gorm:"index;not null"`
	User          User   `gorm:"foreignKey:UserID"`
	Status        string `gorm:"size:20;default:'草稿'"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Comments      []Comment      `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CommentNum    int            `gorm:"default:0"`
	CommentStatus string         `gorm:"size:20;default:'无评论'"`
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"index;not null"`
	User      User   `gorm:"foreignKey:UserID"`
	PostID    uint   `gorm:"index;not null"`
	Post      Post   `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
