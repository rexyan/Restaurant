package main

import (
	"Restaurant/controller"
	"Restaurant/middleware"
	"Restaurant/tool"
	"github.com/gin-gonic/gin"
)

// router
func registerRouter(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
	new(controller.MemberController).Router(engine)
}

// middleware
func registerMiddleware(engine *gin.Engine)  {
	engine.Use(middleware.ResponseMiddleware())
}

// redis
func registerDB()  {
	tool.OrmEngine()
	tool.InitRedisStore()
}

func main() {
	// parse config
	config, _ := tool.ParseConfig("./config/app.json")
	app := gin.Default()
	registerMiddleware(app)
	registerRouter(app)
	registerDB()
	app.Run(config.AppHost + ":" + config.AppPort)
}
