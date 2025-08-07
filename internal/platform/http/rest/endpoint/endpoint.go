package endpoint

import (
	"github.com/charmingruby/pack/pkg/database/postgres"
	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *postgres.Client) {
	api := r.Group("/api")

	v1 := api.Group("/v1")

	v1.GET("/health/live", makeLiveness())
	v1.GET("/health/ready", makeReadiness(db))
}
