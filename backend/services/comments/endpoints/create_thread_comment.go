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

	result := db.Create(&models.ThreadComment{
		ThreadID: uint(id),
		Comment:  body.Body,
		UserID:   uint(userInfo.UserID),
		Likes:    0,
		Comments: 0,
		ParentID: nil,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	result2 := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("comments", gorm.Expr("comments + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	result3 := db.Model(&models.User{}).
		Where("id = ?", uint(userInfo.UserID)).
		Update("comments", gorm.Expr("comments + ?", 1))

	if result3.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	var thread models.Thread
	result4 := db.First(&thread, id)

	if result4.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
	}

	notifications.SendNotification(c, int(thread.UserID), userInfo.Username+" commented on your thread", userInfo.UserID)
	notifications.SendThreadInfo(c, int(thread.ID), "comment", 1)
}
