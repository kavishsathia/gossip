package endpoints

import (
	"backend/helpers"
	"backend/services/threads/usecases"
	"backend/services/threads/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetThread godoc
// @Summary Get a thread
// @Description Get a thread
// @Tags threads
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} thread_types.ThreadResponse "Thread successfully retrieved"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id [get]
func GetThread(c *gin.Context) {
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

	_, userInfo, err := helpers.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !validators.ThreadExists(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "This thread does not exist"})
		return
	}

	usecases.GetThread(c, db, id, userInfo)
}
