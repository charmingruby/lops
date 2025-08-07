package rest

import (
	"github.com/charmingruby/pack/internal/platform/http/rest/endpoint"
	"github.com/charmingruby/pack/pkg/database/postgres"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *postgres.Client,
) {
	endpoint.New(r, db)
}
