package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BodyValidate[T interface{}]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input T

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := c.ShouldBindUri(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := c.ShouldBindQuery(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("body", input)
		c.Next()
	}
}
