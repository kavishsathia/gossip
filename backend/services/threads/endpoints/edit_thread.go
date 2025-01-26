package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EditThread godoc
// @Summary Edits a thread
// @Description Edits a thread
// @Tags threads
// @Accept json
// @Produce json
// @Param thread body thread_types.ThreadCreationForm true "Thread payload"
// @Param id path int true "ID"
// @Success 200 {object} map[string]boolean "Thread successfully edited"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id [put]
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

	// Generate the description
	description, err := helpers.GenerateDescription(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate description"})
		return
	}

	var thread models.Thread
	editThreadResult := db.Model(&thread).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Updates(map[string]interface{}{
			"title":       body.Title,
			"body":        body.Body,
			"image":       body.Image,
			"description": description,
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

	for _, value := range body.Tags {
		db.Create(&models.ThreadTag{
			ThreadID: uint(id),
			Tag:      value,
		})
	}

	c.JSON(http.StatusOK, gin.H{"ThreadID": thread.ID})
}
