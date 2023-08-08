package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Data(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.PureJSON(code, result{Code: code, Message: message, Data: data})
}
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, result{Code: http.StatusOK, Message: "Success", Data: data})
}

func Error(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, result{Code: http.StatusInternalServerError, Message: message})
}

func UnAuth(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, result{Code: http.StatusUnauthorized, Message: message})
}

func NotFound(ctx *gin.Context, message string) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, result{Code: http.StatusNotFound, Message: message})
}
