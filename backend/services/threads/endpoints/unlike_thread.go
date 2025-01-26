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

// UnlikeThread godoc
// @Summary Unlikes a thread
// @Description Unlikes a thread
// @Tags threads
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} map[string]boolean "Thread successfully unliked"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id/like [delete]
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
