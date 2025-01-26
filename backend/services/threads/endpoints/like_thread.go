package endpoints

import (
	"backend/helpers"
	"backend/services/threads/usecases"
	"backend/services/threads/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LikeThread godoc
// @Summary Likes a thread
// @Description Likes a thread
// @Tags threads
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} map[string]boolean "Thread successfully liked"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id/like [post]
func LikeThread(c *gin.Context) {
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

	usecases.LikeThread(c, db, id, userInfo)
}
