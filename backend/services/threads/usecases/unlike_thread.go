package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/notifications"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UnlikeThread(c *gin.Context, db *gorm.DB, id int, userInfo *helpers.User) {
	deleteLikeResult := db.Delete(&models.ThreadLike{
		ThreadID: uint(id),
		UserID:   uint(userInfo.UserID),
	})

	if deleteLikeResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike thread"})
		return
	}

	threadCountUpdateResult := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes - ?", 1))

	if threadCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike thread"})
		return
	}

	notifications.SendThreadInfo(c, id, "like", -1)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
