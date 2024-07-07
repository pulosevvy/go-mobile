package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryValidate[T interface{}]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input T

		if err := c.ShouldBindQuery(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("query", input)
		c.Next()
	}
}
