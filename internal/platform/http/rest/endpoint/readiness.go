package endpoint

import (
	"context"
	"time"

	"github.com/charmingruby/pack/pkg/database/postgres"
	"github.com/charmingruby/pack/pkg/delivery/http/rest"
	"github.com/gin-gonic/gin"
)

const TIMEOUT_IN_SECONDS = 10

func makeReadiness(db *postgres.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT_IN_SECONDS*time.Second)
		defer cancel()

		if err := db.Ping(ctx); err != nil {
			rest.SendServiceUnavailableErrorResponse(c, "database")
			return
		}

		rest.SendOKResponse(c, "", nil)
	}
}
