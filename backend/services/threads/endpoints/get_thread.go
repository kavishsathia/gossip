package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
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
