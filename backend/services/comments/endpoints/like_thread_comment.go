package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/notifications"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LikeThreadComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

	db, err := helpers.OpenDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	_, userInfo, err := helpers.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&models.ThreadCommentLike{
		CommentID: uint(id),
		UserID:    uint(userInfo.UserID),
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	result2 := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	var comment models.ThreadComment
	result3 := db.Table("thread_comments").
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

	if result3.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	notifications.SendNotification(c, int(comment.UserID), userInfo.Username+" liked your comment", userInfo.UserID)

	c.JSON(http.StatusOK, gin.H{"success": true})
}
