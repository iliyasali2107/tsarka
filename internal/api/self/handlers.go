package self_handlers

import (
	"fmt"
	"net/http"

	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
)

type SelfHandlers struct {
	SelfService service.SelfService
}

func NewSelfHandlers(ss service.SelfService) *SelfHandlers {
	return &SelfHandlers{SelfService: ss}
}

func (sh *SelfHandlers) FindHandler(ctx *gin.Context) {
	substr := ctx.Param("substr")
	if substr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect substr"})
		return
	}

	ctxCtx := ctx.Copy().Request.Context()
	identifiers, err := sh.SelfService.Find(ctxCtx, substr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(identifiers) == 0 {
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("there is no identifier that has %s", substr))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"identifiers": identifiers})
}
