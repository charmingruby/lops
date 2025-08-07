package endpoint

import (
	"github.com/charmingruby/lops/pkg/delivery/http/rest"
	"github.com/gin-gonic/gin"
)

func makeReadiness() gin.HandlerFunc {
	return func(c *gin.Context) {
		rest.SendOKResponse(c, "", nil)
	}
}
