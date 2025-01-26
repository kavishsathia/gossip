package usecases

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReportThread(c *gin.Context, db *gorm.DB, id int, userInfo *helpers.User) {
	db.Create(&models.CommunityFlag{
		ThreadID: uint(id),
		UserID:   uint(userInfo.UserID),
	})

	c.JSON(http.StatusOK, gin.H{"success": true})
}
