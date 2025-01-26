package auth

import (
	"backend/helpers"
	"backend/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body UserCreationForm true "User creation payload"
// @Success 200 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 409 {object} map[string]string "Username already exists"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var body UserCreationForm

	if err := c.BindJSON(&body); err != nil {
		return
	}

	db, err := helpers.OpenDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user := &models.User{
		Username:     body.Username,
		PasswordHash: string(hash),
	}

	result := db.Create(user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"UserID": user.ID})
}

// LoginAsUser godoc
// @Summary Login as a user
// @Description Login as a user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body UserCreationForm true "User creation payload"
// @Success 200 {object} map[string]interface{} "Login successfully"
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /user/sign-in [put]
func LoginAsUser(c *gin.Context) {
	var body UserCreationForm

	if err := c.BindJSON(&body); err != nil {
		return
	}

	db, err := helpers.OpenDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}

	var user models.User

	result := db.Where("username = ?", body.Username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "The user does not exist."})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "The password is wrong."})
		return
	}

	token, err := helpers.GenerateJWT(int(user.ID), user.Username)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal Server Error."})
		return
	}

	// Setting the auth token as a cookie on the browser
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "authToken",
		Value:    token,
		MaxAge:   86400,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

// SignOut godoc
// @Summary Sign out
// @Description Sign out
// @Tags auth
// @Accept json
// @Produce json
// @Success 401 {object} map[string]interface{} "Sign out successfully"
// @Router /user/sign-out [get]
func SignOut(c *gin.Context) {

	// Resetting the cookie as en empty string on the browser
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "authToken",
		Value:    "",
		MaxAge:   86400,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Log out successful",
	})
}

// GetMe godoc
// @Summary Get my profile
// @Description Get my profile
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Profile successfully retrieved"
// @Router /user/me [get]
func GetMe(c *gin.Context) {
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

	var userFullInfo ProfileResponse

	result := db.Model(&models.User{}).Select("id, username, created_at, profile_image").First(&userFullInfo, userInfo.UserID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error. Please try again later."})
		}
		return
	}

	c.JSON(http.StatusOK, userFullInfo)
}

// GetMe godoc
// @Summary Get my profile
// @Description Get my profile
// @Tags auth
// @Accept json
// @Produce json
// @Param id path int true "userID"
// @Success 200 {object} map[string]interface{} "Profile successfully retrieved"
// @Router /user/:id [get]
func GetUser(c *gin.Context) {
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

	var userFullInfo ProfileResponse

	result := db.Model(&models.User{}).
		Select("id, username, posts, comments, aura,  created_at, profile_image").
		First(&userFullInfo, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error. Please try again later."})
		}
		return
	}

	c.JSON(http.StatusOK, userFullInfo)
}
