package auth

import (
	"backend/helpers"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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

	token, err := helpers.Generate(int(user.ID), user.Username)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal Server Error."})
		return
	}

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

func SignOut(c *gin.Context) {

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, userFullInfo)
}

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, userFullInfo)
}
