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

func CreateThreadCommentComment(c *gin.Context, db *gorm.DB, id int, body comment_types.CommentCreationForm, userInfo *helpers.User) {
	var comment models.ThreadComment
	db.Where("id = ?", id).First(&comment)

	parentId := uint(id)
	creationResult := db.Create(&models.ThreadComment{
		ThreadID: comment.ThreadID,
		Comment:  body.Body,
		UserID:   uint(userInfo.UserID),
		Likes:    0,
		Comments: 0,
		ParentID: &parentId,
	})

	if creationResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Updating the comment count on thread
	threadCountUpdateResult := db.Model(&models.Thread{}).
		Where("id = ?", comment.ThreadID).
		Update("comments", gorm.Expr("comments + ?", 1))

	if threadCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// Updating the comment count on the parent comment
	commentCountUpdateResult := db.Model(&models.ThreadComment{}).
		Where("id = ?", uint(id)).
		Update("comments", gorm.Expr("comments + ?", 1))

	if commentCountUpdateResult.Error != nil {
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

	var parent models.ThreadComment
	getParentInfoResult := db.Table("thread_comments").
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

	if getParentInfoResult.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parent comment not found"})
		return
	}

	// Sending notifications to the user
	notifications.SendNotification(c, int(parent.UserID), userInfo.Username+" replied to your comment", userInfo.UserID)
	notifications.SendThreadInfo(c, int(comment.ThreadID), "comment", 1)
}
