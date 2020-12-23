package main

import (
	"Restaurant/controller"
	"Restaurant/tool"
	"github.com/gin-gonic/gin"
)

// router
func registerRouter(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
	new(controller.MemberController).Router(engine)
}

func main() {
	config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()
	registerRouter(app)
	app.Run(config.AppHost + ":" + config.AppPort)
}
