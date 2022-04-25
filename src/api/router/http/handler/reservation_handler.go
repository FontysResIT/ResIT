package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
// @Success 200 {array} []model.Reservation "ok"
// @Router /reservation [get]
func (*handler) GetAllReservations(c *gin.Context) {
	var reservations = _logic.GetAllReservations()
	if reservations == nil {
		reservations = make([]model.Reservation, 0)
	}
	c.JSON(http.StatusOK, reservations)
}

func (*handler) GetAllReservationsByDate(c *gin.Context) {
	dateParams := strings.Split(c.Param("date"), "-")
	var intDateParams = []int{}
	for _, v := range dateParams {
		j, err := strconv.Atoi(v)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		intDateParams = append(intDateParams, j)
	}
	date := time.Date(intDateParams[0], time.Month(intDateParams[1]), intDateParams[2], 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60))
	log.Println("RES DATE", date)
	var reservations = _logic.GetAllReservationsByDate(date)
	if reservations == nil {
		reservations = make([]model.Reservation, 0)
	}
	c.JSON(http.StatusOK, reservations)
}

// @Description Create reservation
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Reservation "ok"
// @Router /reservation [post]
func (*handler) CreateReservation(c *gin.Context) {
	var input model.Reservation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservation, err := _logic.CreateReservation(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
		return
	}
	c.JSON(http.StatusOK, reservation)
}

// @Description Cancel reservation
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Reservation "ok"
// @Router /reservation/{id} [put]
func (*handler) UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	query := c.Param("query")
	switch query {
	case "/cancel":
		result, err := _logic.CancelReservation(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			
			return
		}
		
	c.JSON(http.StatusOK, result)
	}
}
