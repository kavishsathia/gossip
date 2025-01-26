package auth_types

type UserCreationForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ProfileResponse struct {
	Username     string
	ID           int
	CreatedAt    string
	Posts        uint
	Comments     uint
	Aura         uint
	ProfileImage string
}
