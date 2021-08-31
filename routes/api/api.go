package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ndaemy/gin-tutorial/routes/api/posts"
)

func AddApiRoutes(rg *gin.RouterGroup) {
	api := rg.Group("/api")
	posts.AddPostRoutes(api)
}
