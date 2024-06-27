package internal

import (
	"applicant-backend/src/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, entities.Error{Message: "method not allowed"})
	})
	r.Use(func(context *gin.Context) {
		context.Next()
		var error string
		for _, err := range context.Errors {
			// log, handle, etc.
			error = err.Error()
		}
		if error != "" {
			context.AbortWithStatusJSON(http.StatusInternalServerError, entities.Error{Message: error})
		}
	})
	return r
}
