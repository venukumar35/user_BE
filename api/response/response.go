package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseApi struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func BadRequestError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, ResponseApi{
		Message: message,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, ResponseApi{
		Message: message,
	})
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, ResponseApi{
		Message: message,
		Data:    data,
	})
}
func UnauthorizedError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, ResponseApi{
		Message: message,
	})
}
