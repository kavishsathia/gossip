package main

import (
	"backend/helpers"
	"backend/models"
	"backend/services/auth"
	"backend/services/notifications"
	"backend/services/threads"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("authToken")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "The token was not found",
			})
			return
		}
		print(2)
		user, err := helpers.Verify(token.Value)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: " + err.Error(),
			})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func main() {
	models.Migrate()
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	r.POST("/user", auth.CreateUser)
	r.PUT("/user", auth.LoginAsUser)

	protected := r.Group("")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/thread", threads.CreateThread)
		protected.GET("/threads", threads.ListThreads)
		protected.GET("/thread/:id", threads.GetThread)
		protected.POST("/thread/:id/like", threads.LikeThread)
		protected.DELETE("/thread/:id/like", threads.UnlikeThread)

		protected.GET("/thread/:id/comments", threads.ListThreadComments)
		protected.POST("/thread/:id/comment", threads.CreateThreadComment)

		protected.GET("/comment/:id/comments", threads.ListThreadCommentComments)
		protected.POST("/comment/:id/comment", threads.CreateThreadCommentComment)

		protected.POST("/comment/:id/like", threads.LikeThreadComment)
		protected.DELETE("/comment/:id/like", threads.UnlikeThreadComment)

		protected.GET("/notifications", notifications.GetNotifications)
		protected.GET("/thread-info/:id", notifications.GetThreadInfo)
		protected.GET("/user", auth.GetMe)
		protected.GET("/user/:id", auth.GetUser)
	}

	r.Run()
}
