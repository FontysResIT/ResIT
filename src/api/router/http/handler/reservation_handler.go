package handler

import (
	"fmt"
	"net/http"

	"github.com/RealSnowKid/ResIT/logic"
	"github.com/RealSnowKid/ResIT/model"
	"github.com/gin-gonic/gin"
)

var _logic logic.IReservationLogic

type handler struct{}

func NewReservationHandler(logic logic.IReservationLogic) *handler {
	_logic = logic
	return &handler{}
}

// @Description Get all reservations
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.Reservation	"ok"
// @Router /reservation [get]
func (*handler) GetAllReservations(c *gin.Context) {
	var reservation = _logic.GetAllReservations()
	c.JSON(http.StatusOK, reservation)
}

func (*handler) CreateReservation(c *gin.Context) {
	fmt.Println("first log")
	var input model.Reservation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(input)
		return
	}
	fmt.Println(input)
	var reservation = _logic.CreateReservation(input)
	c.JSON(http.StatusOK, reservation)

	// var reservationoutcome = _logic.CreateReservation(model.Reservation)
	// c.JSON(http.StatusOK, reservation)
}
