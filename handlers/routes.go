package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	initPingRoute(engine)
}

func initPingRoute(engine *gin.Engine) {

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
