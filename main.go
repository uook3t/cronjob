package main

import (
	"github.com/gin-gonic/gin"
	"github.com/uook3t/base/logger"
	"github.com/uook3t/cronjob/api"
)

func main() {
	r := gin.Default()
	initCommon()

	api.SetupRouter(r)

	_ = r.Run(":8081")
}

func initCommon() {
	logger.Init("cronjob")
}
