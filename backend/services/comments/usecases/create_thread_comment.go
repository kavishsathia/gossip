package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/comments/comment_types"
	"backend/services/notifications"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateThreadComment(c *gin.Context, db *gorm.DB, id int, body comment_types.CommentCreationForm, userInfo *helpers.User) {
	createCommentResult := db.Create(&models.ThreadComment{
		ThreadID: uint(id),
		Comment:  body.Body,
		UserID:   uint(userInfo.UserID),
		Likes:    0,
		Comments: 0,
		ParentID: nil,
	})

	if createCommentResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Updating the comment count on the thread
	threadCountUpdateResult := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("comments", gorm.Expr("comments + ?", 1))

	if threadCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Updating the comment count on the user
	userCountUpdateResult := db.Model(&models.User{}).
		Where("id = ?", uint(userInfo.UserID)).
		Update("comments", gorm.Expr("comments + ?", 1))

	if userCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	var thread models.Thread
	getThreadInfoResult := db.First(&thread, id)

	if getThreadInfoResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Thread not found"})
	}

	// Sending notifications to the user
	notifications.SendNotification(c, int(thread.UserID), userInfo.Username+" commented on your thread", userInfo.UserID)
	notifications.SendThreadInfo(c, int(thread.ID), "comment", 1)
}
