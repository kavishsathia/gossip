package endpoints

import (
	"backend/helpers"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ReportThread godoc
// @Summary Reports a thread
// @Description Reports a thread
// @Tags threads
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} map[string]boolean "Thread successfully reported"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id/report [put]
func ReportThread(c *gin.Context) {
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

	db.Create(&models.CommunityFlag{
		ThreadID: uint(id),
		UserID:   uint(userInfo.UserID),
	})

	c.JSON(http.StatusOK, gin.H{"success": true})
}
