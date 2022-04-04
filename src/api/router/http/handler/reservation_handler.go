package handler

import (
	"net/http"

	"github.com/RealSnowKid/ResIT/logic"
	"github.com/gin-gonic/gin"
)

func GetAllReservations(c *gin.Context) {
	var reservation = logic.GetAllReservations()
	c.JSON(http.StatusOK, reservation)
}