package pkg

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger(debug bool) {
	var err error

	if debug {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatalf("init logger error: %v", err)
	}
}
