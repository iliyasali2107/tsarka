package iin_handlers

import (
	"net/http"

	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
)

type IINHandlers struct {
	IINService service.IINService
}

func NewIINHandlers(iinService service.IINService) *IINHandlers {
	return &IINHandlers{
		IINService: iinService,
	}
}

type IINRequestBody struct {
	Iin string `json:"iin" binding:"required,max=12,min=12"`
}

func (ih *IINHandlers) CheckHandler(ctx *gin.Context) {
	var req IINRequestBody

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "iin format is not correct")
		return
	}

	str, ok := ih.IINService.Check(req.Iin)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"response": "this IIN is not correct"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": str})
}
