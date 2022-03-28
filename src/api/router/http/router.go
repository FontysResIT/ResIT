package httpRouter

import (
	"fmt"

	"github.com/RealSnowKid/ResIT/config"
	"github.com/RealSnowKid/ResIT/router/http/handler"
	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()
	engine := gin.Default()
	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())

	api := engine.Group("/api/")

	//Routes are defined here
	api.GET("/reservation", handler.GetAllReservations)

	fmt.Println(engine.Run(fmt.Sprintf(":%s", config.GetString("http.port"))))
}