package routes

import (
	"github.com/Kaikai20040827/graduation/internal/config"
	"github.com/Kaikai20040827/graduation/internal/handler"
	"github.com/Kaikai20040827/graduation/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册所有 API 路由
// 由 main.go 负责把 handler 和 config 注入进来
func RegisterAPIRoutes(
	r *gin.Engine,
	authH *handler.AuthHandler,
	userH *handler.UserHandler,
	fileH *handler.FileHandler,
	jwtCfg *config.JWTConfig,
) {
	api := r.Group("/api/v1")

	// 公共 API
	{
		api.GET("/ping", handler.Ping)
		
		auth := api.Group("/auth")
		{
			auth.POST("/register", authH.Register)
			auth.POST("/login", authH.Login)
		}
	}

	// 需要认证
	authRequired := api.Group("")
	authRequired.Use(middleware.JWTAuthMiddleware(jwtCfg))
	{
		// 用户
		authRequired.GET("/user/profile", userH.GetProfile)
		authRequired.PUT("/user/profile", userH.UpdateProfile)
		authRequired.PUT("/user/password", userH.ChangePassword)

		// 文件
		authRequired.POST("/files/upload", fileH.UploadFile)
		authRequired.GET("/files", fileH.ListFiles)
		authRequired.GET("/files/download/:id", fileH.DownloadFile)
		authRequired.DELETE("/files/:id", fileH.DeleteFile)
	}
}