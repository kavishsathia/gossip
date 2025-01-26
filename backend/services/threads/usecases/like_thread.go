package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/notifications"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeThread(c *gin.Context, db *gorm.DB, id int, userInfo *helpers.User) {
	likeThreadResult := db.Create(&models.ThreadLike{
		ThreadID: uint(id),
		UserID:   uint(userInfo.UserID),
	})

	if likeThreadResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
		return
	}

	threadCountUpdateResult := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes + ?", 1))

	if threadCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
		return
	}

	var thread models.Thread
	getThreadInfoResult := db.First(&thread, id)

	if getThreadInfoResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
		return
	}

	notifications.SendNotification(c, int(thread.UserID), userInfo.Username+" liked your thread", userInfo.UserID)
	notifications.SendThreadInfo(c, int(thread.ID), "like", 1)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
