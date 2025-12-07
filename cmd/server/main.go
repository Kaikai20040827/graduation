package main

import (
	"log"
	"strconv"

	"github.com/Kaikai20040827/graduation/internal/config"
	"github.com/Kaikai20040827/graduation/internal/handler"
	"github.com/Kaikai20040827/graduation/internal/pkg"
	"github.com/Kaikai20040827/graduation/internal/routes"
	"github.com/Kaikai20040827/graduation/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	Debug = true
)

func main() {
	// 1. 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}

	// 2. Logger
	pkg.InitLogger(Debug)

	// 3. DB
	db, err := pkg.NewDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("db init error: %v", err)
	}

	// 4. Services
	userSrv := service.NewUserService(db)
	fileSrv := service.NewFileService(db, "./storage")

	// 5. Handlers
	authH := handler.NewAuthHandler(userSrv, &cfg.JWT)
	userH := handler.NewUserHandler(userSrv)
	fileH := handler.NewFileHandler(fileSrv)

	// 6. Gin
	r := gin.Default()

	// 7. 注册 API 路由（最关键）
	routes.RegisterAPIRoutes(r, authH, userH, fileH, &cfg.JWT)

	// 8. 启动
	port:= strconv.Itoa(cfg.Server.Port)
	host := cfg.Server.Host
	addr := host + ":" + port
	log.Println("server running at", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("server run error: %v", err)
	}
}
