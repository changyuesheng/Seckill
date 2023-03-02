package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Welcome(ctx *gin.Context) {

	load := ctx.Query("load")
	ctx.String(http.StatusOK, "welcome "+load)
}
