package platform

import (
	"github.com/charmingruby/lops/internal/platform/http/rest"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	rest.New(r)
}
