package comment_types

import "backend/models"

type ThreadCommentResponse struct {
	models.ThreadComment
	Liked        bool
	Username     string
	ProfileImage string
}

type CommentCreationForm struct {
	Body string `json:"body"`
}
