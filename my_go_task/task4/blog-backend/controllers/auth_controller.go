package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my_go_task/task4/blog-backend/config"
	"github.com/my_go_task/task4/blog-backend/models"
	"github.com/my_go_task/task4/blog-backend/utils"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error hashing password", nil)
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error creating user", nil)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, true, "User created successfully", user)
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		utils.JSONResponse(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	if err := user.CheckPassword(input.Password); err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, false, "Invalid password", nil)
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, false, "Error generating token", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, true, "Login successful", gin.H{"token": token})
}