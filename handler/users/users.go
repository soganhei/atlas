package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soganhei.com.br/atlas"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	UsersServices atlas.UsersServices
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
	h.router.POST("/users", h.create)
	h.router.GET("/users", h.find)
}
func (h *Handler) find(c *gin.Context) {

	paginate, err := search(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items, err := h.UsersServices.Find(paginate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
func (h *Handler) create(c *gin.Context) {

	//Pegar token JWT
	//token := c.MustGet("token").(atlas.TokenData)

	var payload atlas.Users

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(*payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pw := string(password)

	payload.Password = &pw

	err = h.UsersServices.Create(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := map[string]interface{}{
		"status": "ok",
	}
	c.JSON(http.StatusCreated, resp)

}
