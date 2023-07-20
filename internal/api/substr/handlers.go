package substr_handlers

import (
	"fmt"
	"net/http"

	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
)

type SubstrHandlers struct {
	SubstrService service.SubstrService
}

func NewSubstHandler(substrService service.SubstrService) *SubstrHandlers {
	return &SubstrHandlers{
		SubstrService: substrService,
	}
}

type FindRequestBody struct {
	Substr string `json:"substr" binding:"required,min=1"`
}

func (sh *SubstrHandlers) FindHandler(ctx *gin.Context) {
	var req FindRequestBody
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "input is not correct"})
		return
	}

	ctxCtx := ctx.Request.Context()

	err = sh.SubstrService.Validate(ctxCtx, req.Substr)
	if err != nil {
		if err == service.ErrMatch {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "input is not correct"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something unexpected happened"})
		return
	}

	substr := sh.SubstrService.Find(ctxCtx, req.Substr)
	ctx.JSON(http.StatusOK, gin.H{fmt.Sprintf("substr of %s", req.Substr): substr})
}
