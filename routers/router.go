package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jalilbengoufa/go-search/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	apiv1.Use()
	{
		apiv1.GET("/word/:id", v1.GetWord)
		apiv1.POST("/word", v1.AddWord)
		apiv1.GET("/words", v1.GetWords)
		apiv1.GET("/search", v1.FindWord)
		apiv1.GET("/autocomplete", v1.Autocomplete)
	}

	return r
}
