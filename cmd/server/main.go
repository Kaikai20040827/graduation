package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Kaikai20040827/graduation/internal/config"
	"github.com/Kaikai20040827/graduation/internal/handler"
	"github.com/Kaikai20040827/graduation/internal/middleware"
	"github.com/Kaikai20040827/graduation/internal/pkg"
	"github.com/Kaikai20040827/graduation/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
    cfg, err := config.LoadConfig()

    if err != nil {
        panic(err)
    }
    pkg.InitLogger(cfg.Server.Debug)

    defer func() {
		if pkg.Logger != nil {
			_ = pkg.Logger.Sync()
		}
	}()

    //db
    db, err := pkg.NewDatabase(&cfg.Database)
	if err!=nil{
		pkg.Logger.Sugar().Fatalf("db connect failed: %v", err)
	
	// services
	userSrv := service.NewUserService(db)
	fileSrv := service.NewFileService(db, "./storage")

}
