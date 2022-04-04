package handler

import (
	"net/http"

	"github.com/RealSnowKid/ResIT/logic"
	"github.com/gin-gonic/gin"
)

var _logic logic.IReservationLogic
type handler struct {}

func NewReservationHandler(logic logic.IReservationLogic) *handler {
	_logic = logic
	return &handler{}
}

func (*handler) GetAllReservations(c *gin.Context) {
	var reservation = _logic.GetAllReservations()
	c.JSON(http.StatusOK, reservation)
}