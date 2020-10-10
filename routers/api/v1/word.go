package v1

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]
func GetWord(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}