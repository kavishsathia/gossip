package usecases

import (
	"backend/models"
	"backend/services/auth/auth_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context, db *gorm.DB, body auth_types.UserCreationForm) {
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user := &models.User{
		Username:     body.Username,
		PasswordHash: string(hash),
	}

	result := db.Create(user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"UserID": user.ID})
}
