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

func UnlikeThread(c *gin.Context) {
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

	result := db.Delete(&models.ThreadLike{
		ThreadID: uint(id),
		UserID:   uint(userInfo.UserID),
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike thread"})
		return
	}

	result2 := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes - ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike thread"})
		return
	}

	notifications.SendThreadInfo(c, id, "like", -1)
	c.JSON(http.StatusOK, gin.H{"success": true})
}
