package pipes

import (
	"api/src/core/validation"
	"api/src/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidationPipe[T comparable]() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := utils.GetBody[T](c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ok, errs := validation.DefaultValidator(body)

		if ok {
			c.Set("body", body)
			c.Next()
			return
		}

		data, err := json.Marshal(errs)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": string(data),
		})
		return
	}
}
