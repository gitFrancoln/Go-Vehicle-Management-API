package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {

			for _, e := range c.Errors {
				logger.Println("Ocurrió un error:", e.Err)
			}

			apiError := c.Errors.Last()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": apiError.Error(),
			})
		}
	}
}
