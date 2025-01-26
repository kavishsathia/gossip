package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetThread(c *gin.Context, db *gorm.DB, id int, userInfo *helpers.User) {
	var thread thread_types.ThreadResponse
	getThreadResult := db.Model(&models.Thread{}).
		Select(`
        threads.id, 
        title, 
        likes, 
        threads.comments, 
        body, 
        description, 
        threads.user_id, 
        shares, 
        threads.created_at, 
        threads.updated_at, 
        threads.image, 
		threads.moderation_flag,
        username, 
		deleted,
        profile_image
    `).
		Joins(`
        INNER JOIN users 
        ON users.id = threads.user_id
    `).
		Preload("ThreadTags").
		Preload("ThreadCorrections").
		First(&thread, id)

	if getThreadResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
		return
	}

	var liked models.ThreadLike
	getThreadLikeResult := db.Where("user_id = ? AND thread_id = ?", userInfo.UserID, id).First(&liked)
	if getThreadLikeResult.Error != nil {
		c.JSON(http.StatusOK, thread_types.ThreadResponse{
			Thread:       thread.Thread,
			Username:     thread.Username,
			ProfileImage: thread.ProfileImage,
			Liked:        false,
		})
		return
	}

	c.JSON(http.StatusOK, thread_types.ThreadResponse{
		Thread:       thread.Thread,
		Username:     thread.Username,
		ProfileImage: thread.ProfileImage,
		Liked:        true,
	})
}
