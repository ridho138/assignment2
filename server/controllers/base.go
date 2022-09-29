package controllers

import (
	"assignment2/server/views"

	"github.com/gin-gonic/gin"
)

func WriteJsonResponse(ctx *gin.Context, payload *views.Response) {
	ctx.JSON(payload.Status, payload)
}
