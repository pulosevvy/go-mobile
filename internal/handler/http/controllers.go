package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	task "go-mobile/internal/handler/http/task"
	user "go-mobile/internal/handler/http/user"
	"go-mobile/internal/service"
	"go-mobile/package/validator/validation"
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

	//Register custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("isUuid", validation.IsUuid)
		if err != nil {
			log.Error("register validator validation error", err)
		}

		err = v.RegisterValidation("passport", validation.Passport)
		if err != nil {
			log.Error("register validator validation error", err)
		}
		err = v.RegisterValidation("dateformat", validation.ValidateDateFormat)
		if err != nil {
			log.Error("register validator validation error", err)
		}
	}
}
