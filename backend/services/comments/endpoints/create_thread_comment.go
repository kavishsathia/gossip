package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/comments/comment_types"
	"backend/services/notifications"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateThreadComment godoc
// @Summary Creates a direct comment
// @Description Creates a direct comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body comment_types.CommentCreationForm true "Comment payload"
// @Param id path int true "threadID"
// @Success 200 {object} map[string]boolean "Comment successfully created"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id/comment [post]
func CreateThreadComment(c *gin.Context) {
	var body comment_types.CommentCreationForm

	if err := c.BindJSON(&body); err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id does not exist or is not an integer"})
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
