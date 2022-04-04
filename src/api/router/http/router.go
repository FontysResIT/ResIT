package httpRouter

import (
	"fmt"
	"net/http"

	"github.com/RealSnowKid/ResIT/config"
	"github.com/RealSnowKid/ResIT/router/http/handler"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Init() {
	config := config.GetConfig()
	engine := gin.Default()
	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())
	engine.Use(static.Serve("/", static.LocalFile("./static", true)))
	api := engine.Group("/api/")
	engine.NoRoute(func(c *gin.Context){
		c.File("./static/index.html")
	})
	//Routes are defined here
	api.GET("/health", healthCheck)
	api.GET("/reservation", handler.GetAllReservations)

	fmt.Println(engine.Run(fmt.Sprintf(":%s", config.GetString("http.port"))))
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "server is up and running"})
}