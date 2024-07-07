package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func UuidValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if _, err := uuid.Parse(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			c.Abort()
			return
		}
		c.Set("id", id)
		c.Next()
	}
}
