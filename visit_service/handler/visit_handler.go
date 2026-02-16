package handler

import (
	"log"
	"net/http"
	"strconv"
	"visit-service/model"
	"visit-service/service"

	"github.com/gin-gonic/gin"
)

type VisitHandler struct {
	service *service.VisitService
}

func NewVisitHandler(s *service.VisitService) *VisitHandler {
	return &VisitHandler{service: s}
}

func (h *VisitHandler) GetVisits(ctx *gin.Context) {
	visits, err := h.service.GetVisits()
	if err != nil {
		log.Printf("GetVisits error: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, visits)
}

func (h *VisitHandler) CreateVisit(ctx *gin.Context) {
	var req model.VisitDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at login request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := h.service.CreateVisit(&req)
	if err != nil {
		log.Printf("Create visit error: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *VisitHandler) DeleteVisit(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.service.DeleteVisit(id)
	if err != nil {
		log.Printf("error of delete visit %s: %s", id, err)
		if _, ok := err.(*strconv.NumError); ok {
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.Status(http.StatusInternalServerError)
		return
	}

}
