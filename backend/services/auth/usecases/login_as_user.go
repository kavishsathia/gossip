package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/auth/auth_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginAsUser(c *gin.Context, db *gorm.DB, body auth_types.UserCreationForm) {
	var user models.User

	result := db.Where("username = ?", body.Username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "The user does not exist."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "The password is wrong."})
		return
	}

	token, err := helpers.GenerateJWT(int(user.ID), user.Username)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal Server Error."})
		return
	}

	// Setting the auth token as a cookie on the browser
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "authToken",
		Value:    token,
		MaxAge:   86400,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}
