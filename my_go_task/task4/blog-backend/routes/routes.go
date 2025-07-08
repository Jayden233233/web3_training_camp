package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/my_go_task/task4/blog-backend/controllers"
	"github.com/my_go_task/task4/blog-backend/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Post routes
	posts := r.Group("/posts")
	{
		posts.GET("", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPost)
		
		// Authenticated routes
		authPosts := posts.Group("")
		authPosts.Use(middlewares.AuthMiddleware())
		{
			authPosts.POST("", controllers.CreatePost)
			authPosts.PUT("/:id", controllers.UpdatePost)
			authPosts.DELETE("/:id", controllers.DeletePost)
		}
	}

	// Comment routes
	comments := r.Group("/posts/:postId/comments")
	{
		comments.GET("", controllers.GetComments)
		
		// Authenticated routes
		authComments := comments.Group("")
		authComments.Use(middlewares.AuthMiddleware())
		{
			authComments.POST("", controllers.CreateComment)
		}
	}

	return r
}