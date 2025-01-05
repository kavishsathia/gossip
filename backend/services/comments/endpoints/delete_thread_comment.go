package endpoints

import (
	"backend/helpers"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteThreadComment(c *gin.Context) {
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

	result := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Update("deleted", true).
		Update("comment", "[deleted]")

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
