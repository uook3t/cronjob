package app

import (
	"github.com/gin-gonic/gin"
	"github.com/uook3t/base/logger"
	"github.com/uook3t/base/response"
)

func Ping(c *gin.Context) {
	ctx := response.GenCtx(c)

	logger.Ctx(ctx).Info("hello world")
	response.Success(c)
}
