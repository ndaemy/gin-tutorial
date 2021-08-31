package posts

import "github.com/gin-gonic/gin"

func AddPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.GET("", list)
	posts.POST("", write)
	posts.GET("/:id", read)
	posts.PATCH("/:id", update)
	posts.DELETE("/:id", remove)
}
