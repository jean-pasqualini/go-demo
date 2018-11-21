package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Printf("Darvador va vers %s...\n", context.Request.URL.String())
	}
}