package middleware

import "github.com/gin-gonic/gin"

func Abort(status int, err error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err != nil {
			c.AbortWithStatusJSON(status, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Next()
	}
}
