package main

import (
	"log"
	"myapp/internal/config"
    "myapp/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	app := routes.SetupRouter(cfg)
	
	app.RunServer(":" + cfg.Server.Port)

	if err := app.Run(":" + cfg.Server.Port); err != nil {
        log.Fatal("服务器启动失败:", err)
    }

}

