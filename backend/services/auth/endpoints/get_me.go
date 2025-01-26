package endpoints

import (
	"backend/helpers"
	"backend/services/auth/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMe godoc
// @Summary Get my profile
// @Description Get my profile
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Profile successfully retrieved"
// @Router /user/me [get]
func GetMe(c *gin.Context) {
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

	usecases.GetMe(c, db, userInfo)
}
