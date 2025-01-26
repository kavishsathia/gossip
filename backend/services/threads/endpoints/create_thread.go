package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateThread godoc
// @Summary Create a new thread
// @Description Creates a new discussion thread
// @Tags threads
// @Accept json
// @Produce json
// @Param thread body thread_types.ThreadCreationForm true "Thread creation payload"
// @Success 200 {object} map[string]interface{} "Thread successfully created"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread [post]
func CreateThread(c *gin.Context) {
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

	// Moderate the thread
	moderationFlag, err := helpers.Moderate(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to moderate"})
		return
	}

	// Generate the description
	description, err := helpers.GenerateDescription(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate description"})
		return
	}

	// Do the fact checking
	corrections, err := helpers.GenerateCorrections(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fact check thread."})
	}

	thread := &models.Thread{
		Title:          body.Title,
		Description:    description,
		Image:          body.Image,
		Body:           body.Body,
		UserID:         uint(userInfo.UserID),
		Likes:          0,
		Comments:       0,
		Shares:         0,
		ModerationFlag: moderationFlag,
	}

	createThreadResult := db.Create(thread)

	if createThreadResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	// Update the user's post count
	userCountUpdateResult := db.Model(&models.User{}).
		Where("id = ?", uint(userInfo.UserID)).
		Update("posts", gorm.Expr("posts + ?", 1))

	if userCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	// Add the tags
	for _, value := range body.Tags {
		db.Create(&models.ThreadTag{
			ThreadID: thread.ID,
			Tag:      value,
		})
	}

	// Add the corrections
	for _, value := range corrections {
		db.Create(&models.ThreadCorrection{
			ThreadID:   thread.ID,
			Correction: value,
		})
	}

	c.JSON(http.StatusOK, gin.H{"ThreadID": thread.ID})
}
