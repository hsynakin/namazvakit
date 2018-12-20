package apiroot

import (
	"github.com/gin-gonic/gin"
	"github.com/hsynakin/namazvakit/controller"
)

func Apiroot(api *gin.RouterGroup) {
	api.GET("/ilceID/:id", controller.Controller)
}
