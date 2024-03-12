package handler

import (
	"github.com/gin-gonic/gin"
	"medods/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/:guid", h.getTokensByGUID)
	router.PUT("/:refresh", h.refreshTokens)

	return router
}
