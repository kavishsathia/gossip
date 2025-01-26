package endpoints

import (
	"backend/helpers"
	"backend/services/auth/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMe godoc
// @Summary Get my profile
// @Description Get my profile
// @Tags auth
// @Accept json
// @Produce json
// @Param id path int true "userID"
// @Success 200 {object} map[string]interface{} "Profile successfully retrieved"
// @Router /user/:id [get]
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

	db, err := helpers.OpenDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	usecases.GetUser(c, db, id)
}
