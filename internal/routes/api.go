package routes

import (
	"github.com/Kaikai20040827/graduation/internal/config"
	"github.com/Kaikai20040827/graduation/internal/handler"
	"github.com/Kaikai20040827/graduation/internal/middleware"
	"github.com/Kaikai20040827/graduation/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine, cfg *config.Config) {
	// services
	userSrv := service.NewUserService(service.DB()) // but we didn't expose DB; we will instead pass gorm.DB directly in main
	// To avoid confusion, main.go will instantiate services and handlers and register routes.
	// This file kept for conceptual grouping but main.go does wiring.
	_ = userSrv
	_ = cfg
}
