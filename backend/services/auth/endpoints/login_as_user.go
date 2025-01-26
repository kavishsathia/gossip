package endpoints

import (
	"backend/helpers"
	"backend/services/auth/auth_types"
	"backend/services/auth/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginAsUser godoc
// @Summary Login as a user
// @Description Login as a user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body UserCreationForm true "User creation payload"
// @Success 200 {object} map[string]interface{} "Login successfully"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/sign-in [put]
func LoginAsUser(c *gin.Context) {
	var body auth_types.UserCreationForm

	if err := c.BindJSON(&body); err != nil {
		return
	}

	db, err := helpers.OpenDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	usecases.LoginAsUser(c, db, body)
}
