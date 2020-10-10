package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jalilbengoufa/go-search/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		
		apiv1.GET("/ping", v1.GetWord)
	}

	return r
}