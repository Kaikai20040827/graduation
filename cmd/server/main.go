package main

import (
	"fmt"
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
	fmt.Println("-----Secure File Box-----")
	fmt.Println("")

	// 1. 加载配置
	fmt.Println("-----Starting loading configuration-----")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}
	fmt.Println("-----Loaded successfully-----")
	fmt.Println("")

	// 2. Logger
	fmt.Println("-----Starting initializing logger-----")
	pkg.InitLogger(Debug)
	fmt.Println("-----Initialized logger successfully-----")
	fmt.Println("")

	// 3. DB(mysql)
	fmt.Println("-----Starting initializing database-----")
	db, err := pkg.NewDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("db init error: %v", err)
	}
	fmt.Println("-----Initialized database successfully-----")
	fmt.Println("")

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
