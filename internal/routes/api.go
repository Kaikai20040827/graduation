package routes

import (
	"github.com/gin-gonic/gin"
)

// 定义运行模式
const (
	DebugMode = "debug"
	ReleaseMode = "release"	
	TestMode = "test"
)

// SetupRouter 设置路由
func SetupRouter(r *gin.Engine) *gin.Engine {
	

	// 全局中间件
	

}
