package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/comments/comment_types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EditThreadComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

	var body comment_types.CommentCreationForm
	if err := c.BindJSON(&body); err != nil {
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
		Update("comment", body.Body)

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
