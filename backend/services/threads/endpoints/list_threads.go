package endpoints

import (
	"backend/helpers"
	"backend/services/threads/usecases"
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

	usecases.ListThreads(c, db, userInfo, tags, people, search)
}
