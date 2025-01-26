package usecases

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListThreads(c *gin.Context, db *gorm.DB, userInfo *helpers.User,
	tags []string, people []string, search string, page int) {
	var threads []thread_types.ThreadResponse

	query := db.Debug().Model(&models.Thread{}).
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
        CASE 
            WHEN utl.user_id IS NOT NULL THEN true 
            ELSE false 
        END as liked
    `).
		Joins(`
        LEFT JOIN thread_likes utl 
        ON utl.thread_id = threads.id 
        AND utl.user_id = ?
    `, userInfo.UserID)

	if len(tags) > 0 {
		print("Hello")
		query = query.
			Joins("LEFT JOIN thread_tags ON thread_tags.thread_id = threads.id").
			Where("thread_tags.tag IN ?", tags)
	}

	if len(people) > 0 {
		query = query.
			Joins("LEFT JOIN users ON users.id = threads.user_id").
			Where("users.username IN ?", people)
	}

	if search != "" {
		query = query.Where("threads.title ILIKE ? OR threads.body ILIKE ?",
			"%"+search+"%",
			"%"+search+"%")
	}

	result := query.
		Limit(10).
		Offset((page - 1) * 10).
		Find(&threads)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch threads"})
		return
	}

	c.JSON(http.StatusOK, threads)
}
