package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Exec CORS Middleware.")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "")
			return
		}
		ctx.Next()
	}
}
