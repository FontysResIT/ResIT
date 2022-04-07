package httpRouter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/RealSnowKid/ResIT/config"
	"github.com/RealSnowKid/ResIT/docs"
	"github.com/RealSnowKid/ResIT/logic"
	"github.com/RealSnowKid/ResIT/repository"
	"github.com/RealSnowKid/ResIT/router/http/handler"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ResIT API
// @version 1.0
// @description The official ResIT API docs
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func Init() {
	config := config.GetConfig()
	engine := gin.Default()
	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())
	engine.Use(static.Serve("/", static.LocalFile("./public", true)))
	api := engine.Group("/api/")

	//Swagger Config & Routes
	docs.SwaggerInfo.BasePath = "/api"
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/api/docs/index.html") })

	//Injection
	reservationRepository := repository.Reservation
	reservationLogic := logic.NewReservationLogic(reservationRepository)
	reservationHandler := handler.NewReservationHandler(reservationLogic)

	engine.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api/") {
			c.File("./public/index.html")
		}
		c.Next()
	})
	//Routes are defined here
	api.GET("/health", healthCheck)
	api.GET("/reservation", reservationHandler.GetAllReservations)
	fmt.Println(engine.Run(fmt.Sprintf(":%s", getPort(config))))
}

// @Description API Healthcheck
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /health [get]
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "server is up and running"})
}

func getPort(config *viper.Viper) string {
	var port string
	if config.GetString("port") != "" {
		port = config.GetString("port")
	} else {
		port = config.GetString("http.port")
	}
	return port
}
