package handler

import (
	"url-shortener/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	publicGroup := router.Group("/auth")
	{
		publicGroup.POST("/sign-in", h.signIn)
		publicGroup.POST("/sign-up", h.signUp)
	}

	privateGroup := router.Group("/storage")
	privateGroup.Use(AuthMiddleware())
	{
		privateGroup.POST("/save", h.saveURL)
		privateGroup.DELETE("/delete", h.deleteURL)
		privateGroup.GET("/get", h.getURL)
	}

	//TODO benchmark Group
	return router
}
