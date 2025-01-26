package endpoints

import (
	"backend/helpers"
	"backend/services/threads/thread_types"
	"backend/services/threads/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
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

	usecases.CreateThread(c, body, userInfo, db)
}
