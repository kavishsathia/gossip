package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
