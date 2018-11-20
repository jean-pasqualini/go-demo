package main

import (
	"github.com/gin-gonic/gin"
	"go-demo/controller"
	_ "net/http"
)

func main() {
	r := gin.Default()
	r.GET("/maison", MaisonController.GetAction)
	r.POST("/maison", MaisonController.PostAction);
	r.Run("0.0.0.0:8181")
}
