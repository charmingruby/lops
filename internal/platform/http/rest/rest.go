package rest

import (
	"github.com/charmingruby/lops/internal/platform/http/rest/endpoint"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	endpoint.New(r)
}
