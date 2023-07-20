package hash_handlers

import (
	"net/http"
	"sync"

	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HashHandlers struct {
	HashService service.HashService
	// TODO:-----------------------
	HashResults sync.Map
}

func NewHashHandlers(hs service.HashService) *HashHandlers {
	return &HashHandlers{
		HashService: hs,
	}
}

func (hh *HashHandlers) RequestsHandler() {
	for {
		req := hh.HashService.GetRequestFromRequests()
		hash := hh.HashService.Hash(req.Str)
		hh.HashResults.Store(req.Id, hash)
	}
}

func (hh *HashHandlers) CalcHandler(ctx *gin.Context) {
	if hh.HashService.IsBusy() {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"reason": "at the current time service is busy, try later"})
		return
	}

	var request struct {
		Str string `json:"str"`
	}

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	id := generateUuid()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ok := hh.HashService.SendRequestToRequests(service.Request{Id: id, Str: request.Str})
	if !ok {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"reason": "at the current time service is busy, try later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"request_id": id})
}

func (hh *HashHandlers) ResultHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	result, ok := hh.HashResults.Load(id)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"message": "PENDING"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"hash": result})
}

func generateUuid() string {
	return uuid.New().String()
}
