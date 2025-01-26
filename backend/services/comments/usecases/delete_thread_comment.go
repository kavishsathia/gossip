package usecases

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteThreadComment(c *gin.Context, db *gorm.DB, id int, userInfo *helpers.User) {
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
