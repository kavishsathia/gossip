package endpoints

import (
	"backend/helpers"
	"backend/services/comments/comment_types"
	"backend/services/comments/usecases"
	"backend/services/comments/validators"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EditThreadComment godoc
// @Summary Edits a comment
// @Description Edits a  comment
// @Tags comments
// @Accept json
// @Produce json
// @Param comment body comment_types.CommentCreationForm true "Comment payload"
// @Param id path int true "commentID"
// @Success 200 {object} map[string]boolean "Comment successfully edited"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /comment/:id [put]
func EditThreadComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

	var body comment_types.CommentCreationForm
	if err := c.BindJSON(&body); err != nil {
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

	if !validators.UserOwnsComment(id, userInfo.UserID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have write access to this comment"})
		return
	}

	usecases.EditThreadComment(c, db, id, userInfo, body)
}
