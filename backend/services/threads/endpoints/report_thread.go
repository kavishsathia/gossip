package endpoints

import (
	"backend/helpers"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
