package platform

import (
	"github.com/charmingruby/pack/internal/platform/http/rest"
	"github.com/charmingruby/pack/pkg/database/postgres"

	"github.com/gin-gonic/gin"
)

func New(
	r *gin.Engine,
	db *postgres.Client,
) {
	rest.New(r, db)
}
