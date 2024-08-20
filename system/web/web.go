package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"steward/api/restful"
	"steward/config"
)

func InitWeb() {
	engine := gin.Default()
	gin.SetMode(config.GetConfig().Base.Model)
	//engine.Use()
	displayBanner()
	restful.LoadRoute(engine)
	if err := engine.Run(fmt.Sprintf(":%s", config.GetConfig().Base.Port)); err != nil {
		panic(fmt.Sprintf("web server start error: %v", err))
	}
}

func displayBanner() {
	f, err := os.ReadFile(config.GetConfig().Base.Banner)
	if err != nil {
		panic(fmt.Sprintf("open banner file error: %v", err))
	}
	fmt.Printf("\n%s\n", f)
}
