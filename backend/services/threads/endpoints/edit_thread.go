package endpoints

import (
	"backend/helpers"
	"backend/services/threads/thread_types"
	"backend/services/threads/usecases"
	"backend/services/threads/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EditThread godoc
// @Summary Edits a thread
// @Description Edits a thread
// @Tags threads
// @Accept json
// @Produce json
// @Param thread body thread_types.ThreadCreationForm true "Thread payload"
// @Param id path int true "ID"
// @Success 200 {object} map[string]boolean "Thread successfully edited"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id [put]
func EditThread(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

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

	if !validators.UserOwnsThread(id, userInfo.UserID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have write access to this thread"})
		return
	}

	usecases.EditThread(c, body, db, id, userInfo)
}
