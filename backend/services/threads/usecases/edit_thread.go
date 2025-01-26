package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EditThread(c *gin.Context, body thread_types.ThreadCreationForm, db *gorm.DB, id int, userInfo *helpers.User) {
	// Generate the description
	description, err := helpers.GenerateDescription(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate description"})
		return
	}

	// Moderate the thread
	moderationFlag, err := helpers.Moderate(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to moderate"})
		return
	}

	// Fact check
	corrections, err := helpers.GenerateCorrections(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fact check thread."})
	}

	var thread models.Thread
	editThreadResult := db.Model(&thread).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Updates(map[string]interface{}{
			"title":           body.Title,
			"body":            body.Body,
			"image":           body.Image,
			"moderation_flag": moderationFlag,
			"description":     description,
		})

	if editThreadResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit thread"})
		return
	}

	if editThreadResult.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to edit thread"})
		return
	}

	// Reset the tags and add them back
	deleteTagsResult := db.Where("thread_id = ?", id).Delete(&models.ThreadTag{})

	if deleteTagsResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit thread"})
		return
	}

	deleteCorrectionsResult := db.Where("thread_id = ?", id).Delete(&models.ThreadCorrection{})

	if deleteCorrectionsResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit thread"})
		return
	}

	for _, value := range body.Tags {
		db.Create(&models.ThreadTag{
			ThreadID: uint(id),
			Tag:      value,
		})
	}

	for _, value := range corrections {
		db.Create(&models.ThreadCorrection{
			ThreadID:   uint(id),
			Correction: value,
		})
	}

	c.JSON(http.StatusOK, gin.H{"ThreadID": thread.ID})
}
