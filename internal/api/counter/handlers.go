package counter_handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
)

type CounterHandlers struct {
	CounterService service.CounterService
}

func NewCounterHandlers(counterService service.CounterService) *CounterHandlers {
	return &CounterHandlers{
		CounterService: counterService,
	}
}

func (ch *CounterHandlers) IncrementHandler(ctx *gin.Context) {
	iStr := ctx.Param("i")
	i, err := strconv.ParseInt(iStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "increment value is incorrect"})
		return
	}

	ctxCtx := ctx.Request.Context()

	res, err := ch.CounterService.Increment(ctxCtx, i)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something unexpected happened: %w", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("counter incremented by %d and now it is %d", i, res)})
}

func (ch *CounterHandlers) DecrementHandler(ctx *gin.Context) {
	iStr := ctx.Param("i")
	i, err := strconv.ParseInt(iStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "decrement value is incorrect"})
		return
	}

	ctxCtx := ctx.Request.Context()

	res, err := ch.CounterService.Decrement(ctxCtx, i)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something unexpected happened: %w", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": fmt.Sprintf("counter decremented by %d and now it is %d", i, res)})
}

func (ch *CounterHandlers) GetValueHandler(ctx *gin.Context) {
	ctxCtx := ctx.Request.Context()
	res, err := ch.CounterService.GetValue(ctxCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("something unexpected happened: %w", err)})
	}

	ctx.JSON(http.StatusOK, gin.H{"value": res})
}
