package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/notifications"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeThreadComment(c *gin.Context, db *gorm.DB, id int, userInfo *helpers.User) {
	createThreadCommentLikeResult := db.Create(&models.ThreadCommentLike{
		CommentID: uint(id),
		UserID:    uint(userInfo.UserID),
	})

	if createThreadCommentLikeResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	// Updating the like count on the comment
	commentCountUpdateResult := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes + ?", 1))

	if commentCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	var comment models.ThreadComment
	getThreadCommentInfoResult := db.Table("thread_comments").
		Select(`
        thread_comments.id,
		thread_comments.user_id, 
		thread_comments.thread_id, 
		thread_comments.comments, 
		thread_comments.likes, 
		thread_comments.comment, 
		thread_comments.parent_id, 
		thread_comments.deleted, 
		thread_comments.created_at, 
		thread_comments.updated_at
    `)

	if getThreadCommentInfoResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	// Sending notification to the user
	notifications.SendNotification(c, int(comment.UserID), userInfo.Username+" liked your comment", userInfo.UserID)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
