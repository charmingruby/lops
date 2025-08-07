package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOKResponse(c *gin.Context, msg string, data any) {
	sendResponse(c, http.StatusOK, msg, data, nil)
}

func sendResponse(c *gin.Context, status int, msg string, data any, err any) {
	isRespEmpty := true

	resp := gin.H{}

	if msg != "" {
		isRespEmpty = false
		resp["message"] = msg
	}

	if data != nil {
		isRespEmpty = false
		resp["data"] = data
	}

	if err != nil {
		isRespEmpty = false
		resp["error"] = err
	}

	if isRespEmpty {
		c.Status(status)
		return
	}

	c.JSON(status, resp)
}
