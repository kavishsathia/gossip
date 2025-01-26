package endpoints

import (
	"backend/helpers"
	"backend/services/comments/comment_types"
	"backend/services/comments/usecases"
	"backend/services/threads/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateThreadComment godoc
// @Summary Creates a direct comment
// @Description Creates a direct comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body comment_types.CommentCreationForm true "Comment payload"
// @Param id path int true "threadID"
// @Success 200 {object} map[string]boolean "Comment successfully created"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /thread/:id/comment [post]
func CreateThreadComment(c *gin.Context) {
	var body comment_types.CommentCreationForm

	if err := c.BindJSON(&body); err != nil {
		return
	}

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

	if !validators.ThreadExists(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "This thread does not exist"})
		return
	}

	usecases.CreateThreadComment(c, db, id, body, userInfo)
}
