package http

import (
	"github.com/gin-gonic/gin"
	"go-mobile/internal/handler/http/error"
	"go-mobile/internal/handler/http/user/dto"
	"go-mobile/internal/service"
	"go-mobile/middleware"
	sl "go-mobile/package/logger/slog"
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
		r.POST("", middleware.BodyValidate[dto.CreateUserDto](), c.Create)
	}
}

func (uc *userController) Create(c *gin.Context) {
	input := c.MustGet("body").(dto.CreateUserDto)

	exists, err := uc.us.GetUserByPassport(c, input.PassportNumber)
	if exists != nil {
		error.NewErrorResponse(c, http.StatusConflict, "user with passport number already exists")
		return
	}
	if err != nil {
		error.InternalServerErrorResponse(c)
		return
	}

	err = uc.us.Create(c, &input)
	if err != nil {
		uc.l.Error("UserController - create", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (uc *userController) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
