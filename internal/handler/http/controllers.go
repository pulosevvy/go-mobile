package http

import (
	"github.com/gin-gonic/gin"
	"go-mobile/internal/handler/http/task"
	"go-mobile/internal/handler/http/user"
	"go-mobile/internal/service"
	"log/slog"
)

func NewControllers(h *gin.Engine, log *slog.Logger, userService service.UserService, taskService service.TaskService) {
	h.Use(gin.Logger())
	h.Use(gin.Recovery())

	routes := h.Group("/api")
	{
		user.NewUserController(routes, log, userService)
		task.NewTaskController(routes, log, taskService)
	}
}
