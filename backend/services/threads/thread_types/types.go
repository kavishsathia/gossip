package thread_types

import "backend/models"

type ThreadCreationForm struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Image string   `json:"image"`
}

type ThreadResponse struct {
	models.Thread
	Liked        bool
	Username     string
	ProfileImage string
}
