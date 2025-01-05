package threads

import (
	"backend/helpers"
	"backend/models"
	"backend/services/notifications"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateThread(c *gin.Context) {
	var body ThreadCreationForm

	if err := c.BindJSON(&body); err != nil {
		print(err)
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

	thread := &models.Thread{
		Title:       body.Title,
		Description: "",
		Image:       body.Image,
		Body:        body.Body,
		UserID:      uint(userInfo.UserID),
		Likes:       0,
		Comments:    0,
		Shares:      0,
	}

	result := db.Create(thread)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	result2 := db.Model(&models.User{}).
		Where("id = ?", uint(userInfo.UserID)).
		Update("posts", gorm.Expr("posts + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	for _, value := range body.Tags {
		db.Create(&models.ThreadTag{
			ThreadID: thread.ID,
			Tag:      value,
		})
	}

	c.JSON(http.StatusOK, gin.H{"ThreadID": thread.ID})
}

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

	var threads []ThreadResponse

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

	var thread ThreadResponse
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
		c.JSON(http.StatusOK, ThreadResponse{
			Thread:       thread.Thread,
			Username:     thread.Username,
			ProfileImage: thread.ProfileImage,
			Liked:        false,
		})
		return
	}

	c.JSON(http.StatusOK, ThreadResponse{
		Thread:       thread.Thread,
		Username:     thread.Username,
		ProfileImage: thread.ProfileImage,
		Liked:        true,
	})
}

func DeleteThread(c *gin.Context) {
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

	result := db.Debug().Model(&models.Thread{}).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Updates(map[string]interface{}{
			"deleted": true,
			"body":    "[deleted]",
		})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func EditThread(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

	var body ThreadCreationForm
	if err := c.BindJSON(&body); err != nil {
		print(err)
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

	var thread models.Thread
	result := db.Model(&thread).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Updates(map[string]interface{}{
			"title": body.Title,
			"body":  body.Body,
			"image": body.Image,
		})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit thread"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to edit thread"})
		return
	}

	result2 := db.Where("thread_id = ?", id).Delete(&models.ThreadTag{})

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit thread"})
		return
	}

	for _, value := range body.Tags {
		db.Create(&models.ThreadTag{
			ThreadID: uint(id),
			Tag:      value,
		})
	}

	c.JSON(http.StatusOK, gin.H{"ThreadID": thread.ID})
}

func LikeThread(c *gin.Context) {
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

	result := db.Create(&models.ThreadLike{
		ThreadID: uint(id),
		UserID:   uint(userInfo.UserID),
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
		return
	}

	result2 := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
		return
	}

	var thread models.Thread
	result3 := db.First(&thread, id)

	if result3.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like thread"})
	}

	notifications.SendNotification(c, int(thread.UserID), userInfo.Username+" liked your thread", userInfo.UserID)
	notifications.SendThreadInfo(c, int(thread.ID), "like", 1)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func UnlikeThread(c *gin.Context) {
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

	result := db.Delete(&models.ThreadLike{
		ThreadID: uint(id),
		UserID:   uint(userInfo.UserID),
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike thread"})
		return
	}

	result2 := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes - ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike thread"})
		return
	}

	notifications.SendThreadInfo(c, id, "like", -1)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func ListThreadComments(c *gin.Context) {
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

	var comments []ThreadCommentResponse
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
		Where("thread_id = ?", id).
		Where("parent_id IS NULL").
		Find(&comments)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

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

	var comments []ThreadCommentResponse
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

func CreateThreadComment(c *gin.Context) {
	var body CommentCreationForm

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

	result := db.Create(&models.ThreadComment{
		ThreadID: uint(id),
		Comment:  body.Body,
		UserID:   uint(userInfo.UserID),
		Likes:    0,
		Comments: 0,
		ParentID: nil,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	result2 := db.Model(&models.Thread{}).
		Where("id = ?", id).
		Update("comments", gorm.Expr("comments + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	result3 := db.Model(&models.User{}).
		Where("id = ?", uint(userInfo.UserID)).
		Update("comments", gorm.Expr("comments + ?", 1))

	if result3.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	var thread models.Thread
	result4 := db.First(&thread, id)

	if result4.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
	}

	notifications.SendNotification(c, int(thread.UserID), userInfo.Username+" commented on your thread", userInfo.UserID)
	notifications.SendThreadInfo(c, int(thread.ID), "comment", 1)
}

func CreateThreadCommentComment(c *gin.Context) {
	var body CommentCreationForm

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

	var comment models.ThreadComment
	db.Where("id = ?", id).First(&comment)
	println(comment.ThreadID)

	parentId := uint(id)
	result := db.Create(&models.ThreadComment{
		ThreadID: comment.ThreadID,
		Comment:  body.Body,
		UserID:   uint(userInfo.UserID),
		Likes:    0,
		Comments: 0,
		ParentID: &parentId,
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	result2 := db.Model(&models.Thread{}).
		Where("id = ?", comment.ThreadID).
		Update("comments", gorm.Expr("comments + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	result3 := db.Model(&models.ThreadComment{}).
		Where("id = ?", uint(id)).
		Update("comments", gorm.Expr("comments + ?", 1))

	if result3.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	result4 := db.Model(&models.User{}).
		Where("id = ?", uint(userInfo.UserID)).
		Update("comments", gorm.Expr("comments + ?", 1))

	if result4.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	var parent models.ThreadComment
	result5 := db.Table("thread_comments").
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
		thread_comments.updated_at
    `)

	if result5.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch threads"})
		return
	}

	notifications.SendNotification(c, int(parent.UserID), userInfo.Username+" replied to your comment", userInfo.UserID)
	notifications.SendThreadInfo(c, int(comment.ThreadID), "comment", 1)
}

func LikeThreadComment(c *gin.Context) {
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

	result := db.Create(&models.ThreadCommentLike{
		CommentID: uint(id),
		UserID:    uint(userInfo.UserID),
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	result2 := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes + ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	var comment models.ThreadComment
	result3 := db.Table("thread_comments").
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
		thread_comments.updated_at
    `)

	if result3.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like comment"})
		return
	}

	notifications.SendNotification(c, int(comment.UserID), userInfo.Username+" liked your comment", userInfo.UserID)

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func UnlikeThreadComment(c *gin.Context) {
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

	result := db.Delete(&models.ThreadCommentLike{
		CommentID: uint(id),
		UserID:    uint(userInfo.UserID),
	})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike comment"})
		return
	}

	result2 := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Update("likes", gorm.Expr("likes - ?", 1))

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func DeleteThreadComment(c *gin.Context) {
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

	result := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Update("deleted", true).
		Update("comment", "[deleted]")

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func EditThreadComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing id or non-integer id"})
		return
	}

	var body CommentCreationForm
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

	result := db.Model(&models.ThreadComment{}).
		Where("id = ?", id).
		Where("user_id = ?", userInfo.UserID).
		Update("comment", body.Body)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
