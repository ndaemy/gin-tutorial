package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ndaemy/gin-tutorial/routes/api"
)

func AddRoutes(e *gin.Engine) {
	rg := e.Group("/")
	api.AddApiRoutes(rg)
}
