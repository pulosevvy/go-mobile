package user

import (
	"github.com/gin-gonic/gin"
	"go-mobile/internal/service"
	"log/slog"
	"net/http"
)

type userController struct {
	us service.UserService
	l  *slog.Logger
}

func NewUserController(route *gin.RouterGroup, l *slog.Logger, us service.UserService) {
	c := &userController{us, l}
	r := route.Group("/users")
	{
		r.GET("", c.GetAll)
		r.POST("", c.Create)
	}
}

func (uc *userController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func (uc *userController) Create(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "OK CREATED",
	})
}
