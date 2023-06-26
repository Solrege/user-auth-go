package main

import (
	"api-rest/cmd/api/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	internal.InitRoutes(r)

	r.Run()
}
