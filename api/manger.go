package api

import (
	"github.com/gin-gonic/gin"
	"github.com/uook3t/cronjob/app"
)

func SetupRouter(r *gin.Engine) {
	r.Use()

	r.GET("/ping/", app.Ping)
	r.POST("")

}
