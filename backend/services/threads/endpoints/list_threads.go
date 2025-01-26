package endpoints

import (
	"backend/helpers"
	"backend/models"
	"backend/services/threads/thread_types"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ListThread godoc
// @Summary Lists threads
// @Description Lists threads
// @Tags threads
// @Accept json
// @Produce json
// @Param query query string true "Search string"
// @Success 200 {object} []thread_types.ThreadResponse "Threads successfully retrieved"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /threads [get]
func ListThreads(c *gin.Context) {
	var queries = strings.Split(c.Query("query"), " ")
	var search = ""
	var tags []string
	var people []string

	for _, query := range queries {
		if len(query) > 0 && query[0] == '#' {
			tags = append(tags, query[1:])
		} else if len(query) > 0 && query[0] == '@' {
			people = append(people, query[1:])
		} else {
			if len(search) > 0 {
				search = search + " " + query
			} else {
				search = query
			}
		}
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
		query = query.Where("threads.title ILIKE ?", "%"+search+"%")
	}

	result := query.Find(&threads)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch threads"})
		return
	}

	c.JSON(http.StatusOK, threads)
}
