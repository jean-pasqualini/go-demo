package main

import (
	"github.com/gin-gonic/gin"
	"go-demo/controller"
	"go-demo/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())
	r.GET("/maison", MaisonController.GetAction)
	r.POST("/maison", MaisonController.PostAction);
	r.Run("0.0.0.0:8181")
}
