package task

import (
	"github.com/gin-gonic/gin"
	"go-mobile/internal/service"
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
		r.GET("", c.GetAll)
		r.POST("", c.Create)
	}
}

func (c *taskController) GetAll(ctx *gin.Context) {

}

func (c *taskController) Create(ctx *gin.Context) {

}
