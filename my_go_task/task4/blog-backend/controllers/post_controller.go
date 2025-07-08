package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my_go_task/task4/blog-backend/config"
	"github.com/my_go_task/task4/blog-backend/models"
	"github.com/my_go_task/task4/blog-backend/utils"
)

func CreatePost(c *gin.Context) {
	userID, err := utils.ExtractTokenID(c)
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}

	post.UserID = userID

	if err := config.DB.Create(&post).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error creating post", nil)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, true, "Post created successfully", post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Preload("User").Preload("Comments").Find(&posts).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error fetching posts", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, "Posts fetched successfully", posts)
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := config.DB.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, false, "Post not found", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, "Post fetched successfully", post)
}

func UpdatePost(c *gin.Context) {
	userID, err := utils.ExtractTokenID(c)
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
		return
	}

	id := c.Param("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, false, "Post not found", nil)
		return
	}

	if post.UserID != userID {
		utils.JSONResponse(c, http.StatusForbidden, false, "You can only update your own posts", nil)
		return
	}

	var updateData models.Post
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}

	if err := config.DB.Model(&post).Updates(updateData).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error updating post", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, "Post updated successfully", post)
}

func DeletePost(c *gin.Context) {
	userID, err := utils.ExtractTokenID(c)
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, false, "Unauthorized", nil)
		return
	}

	id := c.Param("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, false, "Post not found", nil)
		return
	}

	if post.UserID != userID {
		utils.JSONResponse(c, http.StatusForbidden, false, "You can only delete your own posts", nil)
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error deleting post", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, "Post deleted successfully", nil)
}