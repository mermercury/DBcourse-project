package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PaintEgg(ctx *gin.Context) {
	ctx.String(http.StatusOK, "哼～哼～哼～啊啊啊啊啊啊啊啊啊啊～😫😫😫😫😫")
}
