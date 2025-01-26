package endpoints

import (
	"backend/helpers"
	"backend/services/comments/comment_types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListThreadCommentComments godoc
// @Summary Lists nested comment
// @Description Lists nested comment
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "commentID"
// @Success 200 {object} map[string]boolean "Comments succesfully retrieved"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /comment/:id/comments [get]
func ListThreadCommentComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "id does not exist or is not an integer"})
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

	var comments []comment_types.ThreadCommentResponse
	result := db.Table("thread_comments").
		Select(`
        thread_comments.id,
		thread_comments.user_id, 
		thread_comments.thread_id, 
		thread_comments.comments, 
		thread_comments.likes, 
		thread_comments.comment, 
		thread_comments.parent_id, 
		thread_comments.deleted, 
		thread_comments.created_at, 
		thread_comments.updated_at, 
        CASE 
            WHEN utcl.user_id IS NOT NULL THEN true 
            ELSE false 
        END as liked, 
        username, 
        profile_image
    `).
		Joins(`
        INNER JOIN users 
        ON users.id = thread_comments.user_id
    `).
		Joins(`
        LEFT JOIN thread_comment_likes utcl 
        ON utcl.comment_id = thread_comments.id 
        AND utcl.user_id = ?
    `, userInfo.UserID).
		Where("parent_id = ?", id).
		Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
