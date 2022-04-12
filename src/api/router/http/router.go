package httpRouter

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", getIp(config), getPort(config)),
		Handler:      engine,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 20 * time.Second,
	}
	engine.SetTrustedProxies([]string{})
	engine.Use(gin.Recovery())
	engine.Use(static.Serve("/", static.LocalFile("./public", true)))
	api := engine.Group("/api/")
	api.Use(corsMiddleware())
	//Swagger Config & Routes
	docs.SwaggerInfo.BasePath = "/api"
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api.GET("/docs", func(c *gin.Context) { c.Redirect(http.StatusPermanentRedirect, "/api/docs/index.html") })

	//Injection
	reservationRepository := repository.Reservation
	dateTimeSlotRepository := repository.DateTimeSlot
	timeSlotRepository := repository.TimeSlot
	reservationLogic := logic.NewReservationLogic(reservationRepository, dateTimeSlotRepository)
	reservationHandler := handler.NewReservationHandler(reservationLogic)
	dateTimeSlotLogic := logic.NewDateTimeslotLogic(dateTimeSlotRepository)
	dateTimeSlotHandler := handler.NewDateTimeslotHandler(dateTimeSlotLogic)
	timeSlotLogic := logic.NewTimeSlotLogic(timeSlotRepository)
	timeSlotHandler := handler.NewTimeSlotHandler(timeSlotLogic)

	engine.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api/") {
			c.File("./public/index.html")
		}
		c.Next()
	})

	//Routes are defined here
	api.GET("/health", healthCheck)
	api.GET("/reservations", reservationHandler.GetAllReservations)
	// Example: reservations/2022-04-08
	api.GET("/reservations/:date", reservationHandler.GetAllReservationsByDate)
	api.GET("/dateTimeSlots", dateTimeSlotHandler.GetAllDateTimeslots)
	// Example: dateTimeSlots/dateId/2022-04-08
	api.GET("/dateTimeSlots/:query/*param", dateTimeSlotHandler.GetDateTimeslotByParam)
	api.GET("/timeslots", timeSlotHandler.GetAllTimeSlots)
	api.POST("/reservations", reservationHandler.CreateReservation)
	fmt.Printf("[Server] listening on port %s \n", getPort(config))
	s.ListenAndServe()
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

func getIp(config *viper.Viper) string {
	var ip string
	if config.GetString("environment") == "development" {
		ip = "127.0.0.1"
	}
	return ip
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Setting headers")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
