package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

func NewErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Error{message})
	c.Abort()
}

func InternalServerErrorResponse(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Error{"Internal Server Error"})
	c.Abort()
}
