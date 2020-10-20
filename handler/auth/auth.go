package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soganhei.com.br/atlas"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	UsersServices atlas.UsersServices
	JwtServices   atlas.JwtServices
	router        *gin.RouterGroup
}

func NewHandler(router *gin.RouterGroup) *Handler {

	h := Handler{
		router: router,
	}
	h.routers()
	return &h
}

func (h *Handler) routers() {
	h.router.POST("/users/auth", h.auth)

}

func (h *Handler) auth(c *gin.Context) {

	var payload atlas.Users

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenData, err := h.UsersServices.AuthToken(payload.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(tokenData.Password), []byte(*payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.JwtServices.GenerateToken(*tokenData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := map[string]interface{}{
		"token": token,
	}
	c.JSON(http.StatusOK, resp)

}
