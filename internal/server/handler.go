package server

import (
	"regexp/internal/models"
	"regexp/internal/repository"
	"regexp/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine     *gin.Engine
	Services   *services.Services
	Repository *repository.Repository
}

func NewHandler(engine *gin.Engine, services *services.Services) *Handler {
	return &Handler{
		Engine:   engine,
		Services: services,
	}
}

func (h *Handler) Init() {
	h.Engine.GET("/check", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Connected"})
	})
	// h.Engine.POST("/add_card", h.AddUser)
	h.Engine.POST("/add_agent", h.AddAgent)
	h.Engine.POST("/add_service", h.AddService)
	h.Engine.GET("/get_service", h.GetService)
}

func (h *Handler) AddService(ctx *gin.Context) {
	var service *models.CardRule
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Services.AddService(service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Card is added!")
}

func (h *Handler) GetService(ctx *gin.Context) {
	cardNumb := ctx.Query("card_numb")
	log.Println(cardNumb)
	Services, err := h.Services.GetService(cardNumb)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, Services)
}

func (h *Handler) AddAgent(ctx *gin.Context) {
	var agent *models.Agents
	if err := ctx.ShouldBindJSON(&agent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Services.AddAgent(agent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Agent is added!")
}
