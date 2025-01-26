package endpoints

import (
	"backend/services/auth/usecases"

	"github.com/gin-gonic/gin"
)

// SignOut godoc
// @Summary Sign out
// @Description Sign out
// @Tags auth
// @Accept json
// @Produce json
// @Success 401 {object} map[string]interface{} "Sign out successfully"
// @Router /user/sign-out [get]
func SignOut(c *gin.Context) {

	usecases.SignOut(c)
}
