package scripts

import (
	"context"
	"myapp/scripts/logger"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var ctx context.Context
var cancel context.CancelFunc

func Initialize() {
	logger.MakeFile()
	initContext()
	setRouters()
}

func initContext() {
	ctx = context.Background()
	ctx, cancel = context.WithCancel(ctx)
}

func setRouters() {
	router = gin.Default()
	router.GET("/homepage", loginPage)

}

func loginPage(c *gin.Context) {

}

func Run() {
	router.Run(":8080")
}
