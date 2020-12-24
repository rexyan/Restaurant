package main

import (
	"Restaurant/controller"
	"Restaurant/tool"
	"fmt"
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
	if _, err := tool.OrmEngine(config);err!=nil{
		fmt.Println(err.Error())
	}
	app.Run(config.AppHost + ":" + config.AppPort)
}
