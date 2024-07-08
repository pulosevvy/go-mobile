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
		r.GET("", middleware.QueryValidate[dto.GetAllParams](), c.GetAll)
		r.POST("", middleware.BodyValidate[dto.CreateUserDto](), c.Create)
		r.PATCH("/:id", middleware.UuidValidate(), middleware.BodyValidate[dto.UpdateUserDto](), c.Update)
		r.DELETE("/:id", middleware.UuidValidate(), c.Delete)
	}
}

func (uc *userController) GetAll(c *gin.Context) {
	input := c.MustGet("query").(dto.GetAllParams)

	users, err := uc.us.GetAll(c, &input)
	if err != nil {
		uc.l.Error("UserController - GetAll", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusOK, users)
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

	id, err := uc.us.Create(c, &input)
	if err != nil {
		uc.l.Error("UserController - create", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (uc *userController) Update(c *gin.Context) {
	userId := c.MustGet("id").(string)
	input := c.MustGet("body").(dto.UpdateUserDto)

	exists, err := uc.us.GetUserById(c, userId)
	if exists == nil {
		error.NewErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		error.InternalServerErrorResponse(c)
		return
	}
	err = uc.us.Update(c, &input, userId)
	if err != nil {
		uc.l.Error("UserController - Update", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "user updated",
	})
}

func (uc *userController) Delete(c *gin.Context) {
	userId := c.MustGet("id").(string)

	exists, err := uc.us.GetUserById(c, userId)
	if exists == nil {
		error.NewErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		error.InternalServerErrorResponse(c)
		return
	}
	err = uc.us.Delete(c, userId)
	if err != nil {
		uc.l.Error("UserController - Delete", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
