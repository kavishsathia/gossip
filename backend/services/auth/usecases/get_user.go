package usecases

import (
	"backend/models"
	"backend/services/auth/auth_types"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context, db *gorm.DB, id int) {
	var userFullInfo auth_types.ProfileResponse

	result := db.Model(&models.User{}).
		Select("id, username, posts, comments, aura,  created_at, profile_image").
		First(&userFullInfo, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error. Please try again later."})
		}
		return
	}

	c.JSON(http.StatusOK, userFullInfo)
}
