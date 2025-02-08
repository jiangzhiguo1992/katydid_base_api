package router

import (
	"github.com/gin-gonic/gin"
	"katydid_base_api/internal/team/handler"
)

func RegisterV1(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/client/:id", handler.GetClient)
}
