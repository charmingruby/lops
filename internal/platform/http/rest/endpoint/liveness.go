package endpoint

import (
	"github.com/charmingruby/pack/pkg/delivery/http/rest"
	"github.com/gin-gonic/gin"
)

func makeLiveness() gin.HandlerFunc {
	return func(c *gin.Context) {
		rest.SendOKResponse(c, "", nil)
	}
}
