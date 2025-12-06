package main

import (
    "fmt"
    "log"

    "github.com/Kaikai20040827/graduation/internal/config"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("App Name:", cfg.Server.AppName)
    fmt.Println("DB Host:", cfg.Database.Host)

    // 这里启动 Gin / Fiber / Echo 等
}

