package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my_go_task/task4/blog-backend/config"
	"github.com/my_go_task/task4/blog-backend/models"
	"github.com/my_go_task/task4/blog-backend/utils"
)

func CreateComment(c *gin.Context) {
	userID, err := utils.ExtractTokenID(c)
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
		return
	}

	postID := c.Param("postId")
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, false, "Post not found", nil)
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}

	comment.UserID = userID
	comment.PostID = post.ID

	if err := config.DB.Create(&comment).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error creating comment", nil)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, true, "Comment created successfully", comment)
}

func GetComments(c *gin.Context) {
	postID := c.Param("postId")
	var comments []models.Comment
	if err := config.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error fetching comments", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, "Comments fetched successfully", comments)
}