package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	var thread thread_types.ThreadResponse
	result := db.Model(&models.Thread{}).
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
		First(&thread, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
		return
	}

	var liked models.ThreadLike
	result2 := db.Where("user_id = ? AND thread_id = ?", userInfo.UserID, id).First(&liked)
	if result2.Error != nil {
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
