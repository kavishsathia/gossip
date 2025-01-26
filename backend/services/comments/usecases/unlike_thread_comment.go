package usecases

import (
	"backend/helpers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UnlikeThreadComment(c *gin.Context, db *gorm.DB, id int, userInfo *helpers.User) {
	deleteThreadCommentLikeResult := db.Delete(&models.ThreadCommentLike{
		CommentID: uint(id),
		UserID:    uint(userInfo.UserID),
	})

	if deleteThreadCommentLikeResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike comment"})
		return
	}

	// Updating like count on the comment
	commentCountUpdateResult := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes - ?", 1))

	if commentCountUpdateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
