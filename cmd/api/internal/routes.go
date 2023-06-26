package internal

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	h := Handlers{}

	r.GET("/", h.Index)
	r.POST("/register", h.RegisterHandler)
	r.POST("/login", h.LoginHandler)
	r.GET("/homepage", JwtAuthMiddleware(), h.GetPostHandler)

	g1 := r.Group("/post")
	{
		g1.POST("/", JwtAuthMiddleware(), h.NewPostHandler)
		g1.GET("/:id", JwtAuthMiddleware(), h.GetPostByIdHandler)
		g1.DELETE("/:id", JwtAuthMiddleware(), h.DeletePostHandler)
	}

}
