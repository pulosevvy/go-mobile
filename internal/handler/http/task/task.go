package http

import (
	"github.com/gin-gonic/gin"
	"go-mobile/internal/handler/http/error"
	"go-mobile/internal/handler/http/task/dto"
	sl "go-mobile/package/logger/slog"
	"net/http"

	"go-mobile/internal/service"
	"go-mobile/middleware"
	"log/slog"
)

type taskController struct {
	ts service.TaskService
	l  *slog.Logger
}

func NewTaskController(route *gin.RouterGroup, l *slog.Logger, ts service.TaskService) {
	c := &taskController{ts, l}
	r := route.Group("/tasks")
	{
		r.GET("/info/:id", middleware.UuidValidate(), c.GetByUser)
		r.POST("", middleware.BodyValidate[dto.CreateTaskDto](), c.Create)
		r.POST("/start-time/:id", middleware.UuidValidate(), middleware.BodyValidate[dto.StartTaskDto](), c.StartTime)
		r.POST("/end-time/:id", middleware.UuidValidate(), middleware.BodyValidate[dto.EndTaskDto](), c.EndTime)
	}
}

func (tc *taskController) GetByUser(c *gin.Context) {
	userId := c.MustGet("id").(string)

	tasks, err := tc.ts.GetByUserId(c, userId)
	if err != nil {
		tc.l.Error("TaskController - GetByUser", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (tc *taskController) Create(c *gin.Context) {
	input := c.MustGet("body").(dto.CreateTaskDto)

	err := tc.ts.CreateTask(c, &input)
	if err != nil {
		tc.l.Error("UserController - create", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "task created",
	})
}

func (tc *taskController) StartTime(c *gin.Context) {
	taskId := c.MustGet("id").(string)
	input := c.MustGet("body").(dto.StartTaskDto)

	err := tc.ts.StartTime(c, taskId, &input)
	if err != nil {
		tc.l.Error("UserController - StartTime", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "update start time",
	})
}

func (tc *taskController) EndTime(c *gin.Context) {
	taskId := c.MustGet("id").(string)
	input := c.MustGet("body").(dto.EndTaskDto)

	exists, err := tc.ts.GetTaskById(c, taskId)
	if exists == nil {
		error.NewErrorResponse(c, http.StatusNotFound, "task not found")
		return
	}
	if exists.StartTask == nil {
		error.NewErrorResponse(c, http.StatusNotFound, "task not found")
		return
	}
	if err != nil {
		error.InternalServerErrorResponse(c)
		return
	}

	err = tc.ts.EndTime(c, exists, &input)
	if err != nil {
		tc.l.Error("TaskController - EndTime", sl.Err(err))
		error.InternalServerErrorResponse(c)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "update end time",
	})
}
