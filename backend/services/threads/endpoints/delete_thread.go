package endpoints

import (
	"backend/helpers"
	"backend/services/threads/usecases"
	"backend/services/threads/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteThread godoc
// @Summary Deletes a thread
// @Description Deletes a thread
// @Tags threads
// @Accept json
// @Produce json
// @Param id path int true "threadID"
// @Success 200 {object} map[string]boolean "Thread successfully deleted"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id [delete]
func DeleteThread(c *gin.Context) {
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

	if !validators.UserOwnsThread(id, userInfo.UserID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have write access to this thread"})
		return
	}

	usecases.DeleteThread(c, db, id, userInfo)
}
