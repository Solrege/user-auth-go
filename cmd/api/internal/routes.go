package internal

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	h := Handlers{}

	r.GET("/", h.Index)
	r.POST("/register", h.RegisterHandler)
	r.POST("/login", h.LoginHandler)
	r.GET("/homepage", JwtAuthMiddleware(), h.GetPostHandler)

	g1 := r.Group("/post", JwtAuthMiddleware())
	{
		g1.POST("/", h.NewPostHandler)
		g1.GET("/:id", h.GetPostByIdHandler)
		g1.DELETE("/:id", h.DeletePostHandler)
		g1.GET("/:id/comments", h.GetCommentsHandler)
		g1.POST("/:id/comments", h.NewCommentHandler)
		g1.DELETE("/:id/comments/:commentId", h.DeleteCommentHandler)
	}

}
