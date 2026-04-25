package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorWrapper(handler func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
		}
	}
}
