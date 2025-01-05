package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EditThread(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

	var body thread_types.ThreadCreationForm
	if err := c.BindJSON(&body); err != nil {
		print(err)
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

	var thread models.Thread
	result := db.Model(&thread).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Updates(map[string]interface{}{
			"title": body.Title,
			"body":  body.Body,
			"image": body.Image,
		})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit thread"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to edit thread"})
		return
	}

	result2 := db.Where("thread_id = ?", id).Delete(&models.ThreadTag{})

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit thread"})
		return
	}

	for _, value := range body.Tags {
		db.Create(&models.ThreadTag{
			ThreadID: uint(id),
			Tag:      value,
		})
	}

	c.JSON(http.StatusOK, gin.H{"ThreadID": thread.ID})
}
