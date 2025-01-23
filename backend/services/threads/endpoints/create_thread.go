package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	flag, err := helpers.Moderate(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to moderate"})
		return
	}

	description, err := helpers.GenerateDescription(c, body.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate description"})
		return
	}

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
		ModerationFlag: flag,
	}

	result := db.Create(thread)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	result2 := db.Model(&models.User{}).
		Where("id = ?", uint(userInfo.UserID)).
		Update("posts", gorm.Expr("posts + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	for _, value := range body.Tags {
		db.Create(&models.ThreadTag{
			ThreadID: thread.ID,
			Tag:      value,
		})
	}

	for _, value := range corrections {
		db.Create(&models.ThreadCorrection{
			ThreadID:   thread.ID,
			Correction: value,
		})
	}

	c.JSON(http.StatusOK, gin.H{"ThreadID": thread.ID})
}
