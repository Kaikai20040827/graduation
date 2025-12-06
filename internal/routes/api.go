package routes

import (
	"github.com/gin-gonic/gin"
)

func RunServer(app *gin.Engine, port string) error {
	return app.Run(":" + port)
}
